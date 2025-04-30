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
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o calculator-swagger ./cmd/calculator_swagger

# Run stage
FROM alpine:latest

# Install ca-certificates for HTTPS connections if needed
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/calculator-swagger /

# Create directory for swagger files
RUN mkdir -p /root/cmd/calculator_swagger

# Copy the swagger.json file
COPY --from=builder /app/cmd/calculator_swagger/swagger.json /swagger.json
# Expose the Swagger UI port
EXPOSE 8082

# Set environment variable for port (can be overridden at runtime)
ENV SWG_PORT=8082

# Run the application
CMD ["/calculator-swagger"]