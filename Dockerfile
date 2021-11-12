FROM golang:latest
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o mongorpc cmd/mongorpc/main.go
EXPOSE 9090
CMD [ "/app/mongorpc" ]