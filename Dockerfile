FROM golang:1.6

RUN mkdir /app

ADD . /app/gonic

WORKDIR /app/gonic

EXPOSE 8080

RUN go get github.com/gin-gonic/gin

RUN go get gopkg.in/mgo.v2
RUN go get gopkg.in/mgo.v2/bson
RUN go build -o app




ENTRYPOINT ["/app/gonic/app"]

