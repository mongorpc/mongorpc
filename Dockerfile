FROM golang:1.14-alpine3.11
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o mongorpc cmd/mongorpc/main.go
EXPOSE 9090
CMD [ "/app/mongorpc" ]