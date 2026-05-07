package gRPC_app

import (
	"context"
	"errors"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/labstack/echo/v4"
	"github.com/tqhuy-dev/gore/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type HealthService struct {
	pb.UnimplementedHealthCheckServiceServer
}

func (s *HealthService) ReadinessCheck(ctx context.Context, req *pb.ReadinessRequest) (*pb.ReadinessResponse, error) {
	return &pb.ReadinessResponse{Version: "1"}, nil
}

func (s *HealthService) SamplingRequest(ctx context.Context, req *pb.SampRequest) (*pb.SamplingReturn, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return &pb.SamplingReturn{Data: req.Data}, nil
}

func StartGRPCServer(grpcPort string, stopCh <-chan struct{}) error {
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	defer func() {
		fmt.Println("grpc server stopped")
	}()

	s := grpc.NewServer()
	pb.RegisterHealthCheckServiceServer(s, &HealthService{})

	// Start gRPC server in goroutine
	go func() {
		fmt.Println("gRPC listening on", grpcPort)
		if err := s.Serve(lis); err != nil {
			fmt.Println("gRPC Server Error:", err)
		}
	}()

	<-stopCh
	time.Sleep(5 * time.Second)
	fmt.Println("Gracefully stopping gRPC server...")
	s.GracefulStop()
	return nil
}

func StartEchoServer(httpPort string, grpcAddr string, stopCh <-chan struct{}) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer func() {
		fmt.Println("Echo server stopped")
		cancel()
	}()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := pb.RegisterHealthCheckServiceHandlerFromEndpoint(ctx, mux, grpcAddr, opts)
	if err != nil {
		return err
	}

	e := echo.New()

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	e.Any("/*", echo.WrapHandler(mux))

	go func() {
		fmt.Println("Echo HTTP API listening on", httpPort)
		if err := e.Start(httpPort); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Println("HTTP server error:", err)
		}
	}()

	<-stopCh
	time.Sleep(5 * time.Second)
	fmt.Println("Gracefully shutting down HTTP server...")
	return e.Shutdown(nil)
}

func RunApp() {
	grpcAddr := ":50051"
	httpAddr := ":8080"

	stopCh := make(chan struct{})
	sigCh := make(chan os.Signal, 1)

	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	// Start HTTP and gRPC in goroutines
	go func() {
		if err := StartEchoServer(httpAddr, grpcAddr, stopCh); err != nil {
			panic(err)
		}
	}()

	go func() {
		if err := StartGRPCServer(grpcAddr, stopCh); err != nil {
			panic(err)
		}
	}()

	// Wait for termination signal
	<-sigCh
	close(stopCh)

	// Delay 2s để cho phép shutdown
	time.Sleep(10 * time.Second)
	fmt.Println("Gracefully shutting down all server...")
}
