
FROM golang:1.12.4

# Set GOPATH/GOROOT environment variables
RUN mkdir -p /go
ENV GOPATH /go
ENV PATH $GOPATH/bin:$PATH

# go get all of the dependencies
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# Set up app
WORKDIR /go/src/service-demo
COPY . .
RUN dep init
RUN dep ensure
RUN go build

EXPOSE 12001

CMD ["./service-demo"]