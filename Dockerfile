# Start from the official Go language base image
FROM golang:1.17-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files
COPY go.mod go.sum ./

# Download and cache Go modules
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o myapp

# Set the entry point for the container
CMD ["./myapp"]
