FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY cmd ./cmd
COPY config ./config
COPY internal ./internal
COPY .env ./.env

RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/app/main.go

CMD ./app