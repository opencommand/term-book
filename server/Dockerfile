FROM golang:1.24 as builder

WORKDIR /app

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code and build the application
COPY . .
RUN go build -o main .

# Use a minimal runtime image
FROM debian:bullseye-slim

WORKDIR /app

# Copy the binary and config file into the image
COPY --from=builder /app/main .
COPY boot.yaml .

# Set the default config path
ENV CONFIG_PATH=/app/boot.yaml

# Run the application using the default config
CMD ["sh", "-c", "./main"]
