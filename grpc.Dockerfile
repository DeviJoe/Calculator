# Build stage
FROM golang:1.24-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o calculator-grpc ./cmd/calculator_grpc

# Run stage
FROM alpine:latest

# Install ca-certificates for HTTPS connections if needed
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/calculator-grpc .

# Expose the gRPC port
EXPOSE 8081

# Set environment variable for port (can be overridden at runtime)
ENV CALCG_PORT=8081

# Run the application
CMD ["./calculator-grpc"]