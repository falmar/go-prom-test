FROM golang:1.9-alpine

ENV PORT 80
EXPOSE 8080

RUN mkdir /go-app
WORKDIR /go-app

COPY prom /go-app

CMD ["./prom"]
