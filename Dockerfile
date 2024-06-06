# Use the official Golang image to create a build artifact.
FROM golang:1.20 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main ./cmd/main.go

# Start a new stage from alpine
FROM alpine:latest

RUN apk --no-cache add ca-certificates

# Install the required dependencies for Go binaries (if any)
RUN apk --no-cache add libc6-compat

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Ensure the binary has executable permissions
RUN chmod +x /root/main

# Command to run the executable
CMD ["./main"]
