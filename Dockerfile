
FROM golang:1.12.4-alpine3.9

RUN apk --no-cache add curl
RUN apk --no-cache add git
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# Set up app
WORKDIR /go/src/service-demo
COPY . .

RUN dep init 
RUN dep ensure
RUN go install -v ./...

EXPOSE 12001

CMD ["service-demo"]