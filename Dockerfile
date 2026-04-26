# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY app/go.mod ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY app/ .

# Build the application
RUN go build -o stress-test ./cmd

# Final stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

# Copy the binary from builder stage
COPY --from=builder /app/stress-test /usr/local/bin/stress-test

# Make sure the binary is executable
RUN chmod +x /usr/local/bin/stress-test

# Set the entrypoint
ENTRYPOINT ["/usr/local/bin/stress-test"]