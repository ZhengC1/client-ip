FROM golang:alpine

ADD ./src client_ip/src/app

WORKDIR /go/src/app

ENV PORT=3001

CMD ["go", "run", "check_ip.go"]

