FROM golang:alpine3.13

WORKDIR /go/src/app
COPY ./webapp.go ./webapp.go

RUN go build webapp.go

EXPOSE 8923

CMD ["./webapp"]

