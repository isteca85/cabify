FROM ubuntu:22.04

WORKDIR /usr/src/app

COPY . .

RUN apt-get update

RUN apt-get install -y ca-certificates

RUN apt-get install -y golang-go

RUN go mod tidy

EXPOSE 3909
