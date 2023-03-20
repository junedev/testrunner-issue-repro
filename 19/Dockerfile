FROM golang:1.19.7-alpine3.17

RUN adduser --disabled-password appuser

COPY . .

WORKDIR app
RUN go build -o main

USER appuser

CMD ["./main"]