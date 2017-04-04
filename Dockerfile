FROM golang:1.8-alpine

MAINTAINER Andrew Mason <andrew@fixedpoint.xyz>

EXPOSE 8080

WORKDIR /code
ADD . /code

RUN go build -o main .

CMD /code/main
