
FROM golang:1.12.4-alpine3.9

RUN apk --no-cache add git

WORKDIR /go/src/service-demo
COPY . .

RUN go get -v ./...
RUN go install -v ./...

EXPOSE 12001

CMD ["service-demo"]