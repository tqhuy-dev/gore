# Step 1: Build Stage
FROM golang:1.23-alpine AS builder

# Set working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum from the root directory (shared for multiple apps)
COPY go.mod go.sum ./
#COPY private.pem public.pem ./
# Download dependencies before copying the application code
RUN go mod download

# Copy only the application source code from the application/ folder
COPY . .
# Set the working directory to the specific application
WORKDIR /app/

# Build the Go binary with optimizations
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o main .

# Step 2: Minimal Runtime Image
#FROM gcr.io/distroless/static:latest
FROM alpine:latest

# Cài shell và công cụ debug nếu cần
RUN apk add --no-cache bash curl

# Set the working directory
WORKDIR /root/

# Copy only the compiled binary from the build stage
COPY --from=builder /app/main .
#COPY --from=builder /app/private.pem ./
#COPY --from=builder /app/public.pem ./
# Expose ports for the application and Prometheus metrics

#ENV GOOGLE_APPLICATION_CREDENTIALS="/root/gcpkms-464615-0903c13490d1.json"

EXPOSE 1323

# Set the entrypoint to run the app
ENTRYPOINT ["/root/main"]
