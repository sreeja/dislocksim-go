FROM golang:1.15.6


RUN mkdir -p /usr/app
RUN mkdir -p /usr/data

WORKDIR /usr/app
COPY . .

RUN go get -d -v ./... 
RUN go build
RUN go install
