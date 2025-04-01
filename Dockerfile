# Use Go as a build stage
FROM golang:1.24 AS builder

WORKDIR /app
COPY . .

# Build a fully static binary
RUN go mod tidy && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o server

# Use a minimal runtime image (Alpine is causing issues, so switch to Debian)
FROM debian:stable-slim

WORKDIR /root/
COPY --from=builder /app/server .

# Ensure the binary is executable
RUN chmod +x /root/server

EXPOSE 8080

CMD ["./server"]

