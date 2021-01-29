FROM golang:1.15.6

RUN apt-get update && apt-get install -y iproute2 iputils-ping

RUN mkdir -p /usr/app
RUN mkdir -p /usr/data

WORKDIR /usr/app
COPY . .

RUN go get -d -v ./... 
RUN go build
RUN go install
