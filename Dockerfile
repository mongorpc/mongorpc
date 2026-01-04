FROM golang:1.22-alpine AS builder

WORKDIR /app

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /mongorpc ./cmd/server

# Runtime image
FROM alpine:3.20

RUN apk --no-cache add ca-certificates

COPY --from=builder /mongorpc /usr/local/bin/mongorpc

EXPOSE 50051

ENTRYPOINT ["mongorpc"]
CMD ["--address", "0.0.0.0:50051"]