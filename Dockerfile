FROM golang:1.6

RUN go get github.com/gin-gonic/gin

EXPOSE 8080