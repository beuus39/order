FROM golang:1.12-alpine

ARG ENVIRONMENT=DEV
ARG DB_NAME=orders
ARG DB_HOST=192.168.1.7
ARG DB_USRERNAME=beu
ARG DB_PASSWORD=Beu$2021
ARG DB_DIALECT=postgres
ARG DB_PORT=5432

ENV ENVIRONMENT=${ENVIRONMENT}
ENV DB_NAME=${DB_NAME}
ENV DB_HOST=${DB_HOST}
ENV DB_USRERNAME=${DB_USRERNAME}
ENV DB_PASSWORD=${DB_PASSWORD}
ENV DB_DIALECT=${DB_DIALECT}
ENV DB_PORT=${DB_PORT}

RUN apk add --no-cache git

RUN mkdir /app
COPY . /app

WORKDIR /app
RUN go mod download

WORKDIR /app/cmd

RUN go build -o main .
CMD ["/app/cmd/main"]