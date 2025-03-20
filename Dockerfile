# Step 1: Build Stage
FROM golang:1.23-alpine AS builder

# Set working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum from the root directory (shared for multiple apps)
COPY go.mod go.sum ./

# Download dependencies before copying the application code
RUN go mod download

# Copy only the application source code from the application/ folder
COPY . .

# Set the working directory to the specific application
WORKDIR /app/example/redis-sentinel-app

# Build the Go binary with optimizations
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o main .

# Step 2: Minimal Runtime Image
FROM gcr.io/distroless/static:latest

# Set the working directory
WORKDIR /root/

# Copy only the compiled binary from the build stage
COPY --from=builder /app/example/redis-sentinel-app/main .

# Expose ports for the application and Prometheus metrics
EXPOSE 8081 2223

# Set the entrypoint to run the app
ENTRYPOINT ["/root/main"]
