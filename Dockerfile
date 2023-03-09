FROM golang:latest

WORKDIR /go/src/app

COPY . .

RUN go mod tidy

RUN go build -o crud-api cmd/crud-api/main.go

EXPOSE 8080

CMD ["./crud-api"]