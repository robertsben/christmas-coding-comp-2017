# Example Dockerfile, edit to your requirements
FROM golang:1-alpine

WORKDIR /go/src/app

COPY src .

RUN go-wrapper download

RUN go-wrapper install

CMD ["go-wrapper", "run"]

