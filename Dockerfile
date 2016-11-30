FROM golang:1.6

RUN go get github.com/gin-gonic/gin

EXPOSE 8080

RUN go build -o app

RUN ./app