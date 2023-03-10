FROM golang:latest

WORKDIR /app

COPY ./ /app/

## Modules Dependency

# .env Manager
RUN go get github.com/spf13/viper

## Dev Environment Dependencies

# Hot Reloading
RUN go install github.com/cosmtrek/air@latest 
RUN export PATH=$PATH:$(go env GOPATH)/bin

COPY go.mod go.sum ./
RUN go mod download
RUN go mod tidy

RUN go build -o crud-api cmd/crud-api/main.go

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]