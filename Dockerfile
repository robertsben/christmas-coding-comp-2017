FROM golang:1-alpine

RUN mkdir /app

WORKDIR /app

ADD src/ .

RUN go build -o main .

CMD ["./main"]