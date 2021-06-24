FROM golang:1.12-alpine

RUN apk add --no-cache git

RUN mkdir /app
COPY . /app

WORKDIR /app
RUN go mod download

WORKDIR /app/cmd

RUN go build -o main .
CMD ["/app/cmd/main"]