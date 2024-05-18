# Use the Go 1.22 official image as the build environment
FROM golang:1.22 AS builder

# Disable CGO
ENV CGO_ENABLED=0

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code and build the app
COPY . .
RUN go build -ldflags "-s -w" -o /app/groqai2api .

# Use Alpine Linux as the final image
FROM alpine:latest

# Set working directory
WORKDIR /app

# Copy the compiled application and resources from the build phase
COPY --from=builder /app/groqai2api /app/groqai2api

# Exposed port
EXPOSE 8080

CMD ["/app/groqai2api"]
