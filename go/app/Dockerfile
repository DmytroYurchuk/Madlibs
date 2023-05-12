FROM golang:1.17.3-alpine

RUN apk update && apk add git && apk add build-base

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go test

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
