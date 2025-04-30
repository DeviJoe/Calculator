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
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o calculator-rest ./cmd/calculator_rest

# Run stage
FROM alpine:latest

# Install ca-certificates for HTTPS connections if needed
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/calculator-rest .

# Expose the REST API port
EXPOSE 8080

# Set environment variable for port (can be overridden at runtime)
ENV CALCR_PORT=8080

# Run the application
CMD ["./calculator-rest"]