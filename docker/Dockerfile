FROM golang:1.20.2

ENV APP_STAGE $APP_STAGE

WORKDIR /app

COPY go.mod ./
COPY ./ /app/

RUN go mod download
RUN go mod tidy

RUN if [ "$APP_STAGE" = "local" ]; \
    then \
        go install github.com/cosmtrek/air@latest && \
        export PATH=$PATH:$(go env GOPATH)/bin; \
fi

RUN go build -o ./bin/${APP_NAME} ./cmd/server/main.go;

EXPOSE 8080 

ENTRYPOINT ["./bin/${APP_NAME}"]