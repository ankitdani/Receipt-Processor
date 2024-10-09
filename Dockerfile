# Start with the official Golang image to build the Go app
FROM golang:1.20-alpine AS builder

# Set the working directory in the container
WORKDIR /app

# Copy the Go modules and dependencies first to build them before copying other files
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy the entire source code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Use a minimal base image for the final container
FROM alpine:latest

# Set the working directory in the final image
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Expose port 8080
EXPOSE 8080

# Command to run the Go app
CMD ["./main"]
