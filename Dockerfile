FROM golang:1.23

COPY . /app
WORKDIR /app

RUN go mod download

RUN go build "cmd/master/main.go"

CMD ["./main"]