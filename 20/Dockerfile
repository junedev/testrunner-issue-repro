FROM golang:1.20.2-alpine3.17

RUN adduser --disabled-password --gecos "" appuser

COPY . .

WORKDIR app
RUN go build -o main

USER appuser

CMD ["./main"]