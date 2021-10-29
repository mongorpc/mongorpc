FROM golang:alpine AS builder
WORKDIR /mongorpc
ADD . .
RUN go mod tidy
RUN go build -o /mongorpc cmd/mongorpc/main.go

FROM alpine:latest
ARG MONGO_URI
ENV MONGO_URI=$MONGO_URI
COPY --from=builder /mongorpc /
EXPOSE 50051
CMD ["mongorpc"]