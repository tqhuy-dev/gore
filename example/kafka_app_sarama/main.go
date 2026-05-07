package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/tqhuy-dev/gore/pb"
	"google.golang.org/protobuf/proto"
	"log"
	"os"
	"os/signal"

	"github.com/IBM/sarama"
)

func main() {
	brokers := []string{"localhost:29092"}
	topic := "jaeger-spans"
	groupID := "jaeger-debug-group"

	config := sarama.NewConfig()
	config.Version = sarama.V2_5_0_0
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	consumerGroup, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		log.Fatalf("Error creating consumer group: %v", err)
	}
	defer consumerGroup.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	handler := &jaegerConsumer{}

	go func() {
		for {
			if err := consumerGroup.Consume(ctx, []string{topic}, handler); err != nil {
				log.Printf("Error from consumer: %v", err)
			}
		}
	}()

	log.Println("Kafka consumer started. Ctrl+C to exit.")
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, os.Interrupt)
	<-sigterm
	log.Println("Shutting down...")
}

// jaegerConsumer implements sarama.ConsumerGroupHandler
type jaegerConsumer struct{}

func (c *jaegerConsumer) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (c *jaegerConsumer) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (c *jaegerConsumer) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		span := &pb.Span{}
		err := proto.Unmarshal(msg.Value, span)
		if err != nil {
			log.Printf("Failed to unmarshal Protobuf span: %v", err)
			continue
		}
		var parentSpanId []byte
		for _, child := range span.References {
			if child.RefType == pb.SpanRefType_CHILD_OF {
				parentSpanId = child.SpanId
			}
		}
		fmt.Printf("Service=%s TraceID=%s Operation=%s SpanID=%s ParentSpanID=%s  Start=%d Duration=%d\n Tags=%+v\n",
			span.Process.ServiceName,
			base64.StdEncoding.EncodeToString(span.TraceId),
			span.OperationName,
			base64.StdEncoding.EncodeToString(span.SpanId),
			base64.StdEncoding.EncodeToString(parentSpanId),
			span.StartTime.GetNanos(),
			span.Duration.GetNanos(),
			span.Tags,
		)
	}
	return nil
}
