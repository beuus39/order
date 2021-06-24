FROM golang:1.10-alpine

RUN mkdir /app
COPY . /app
WORKDIR /app

RUN go build -o main .
CMD ["/app/cmd/main"]