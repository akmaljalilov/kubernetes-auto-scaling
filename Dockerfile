FROM golang:1.17-alpine3.14 as builder
ARG BUILD_VERSION=1
ENV VERSION=$BUILD_VERSION
COPY . /server
WORKDIR /server
RUN go mod download all
RUN CGO_ENABLED=0 GOOS=linux go build -o test-server -a -installsuffix cgo main.go
CMD ["./test-server"]
