
FROM golang:1.12.4-alpine3.9

RUN apk --no-cache add git

WORKDIR $GOPATH/src/gitlab.com/chinnawat.w/golang-service-demo
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 12001

CMD ["golang-service-demo"]