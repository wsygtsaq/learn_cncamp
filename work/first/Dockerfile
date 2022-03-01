FROM golang:1.17-alpine
ENV GOPATH=/go

WORKDIR /go/src/test/first
RUN mkdir src pkg bin
COPY . .
RUN go mod init
RUN go mod tidy
RUN go build
ENTRYPOINT ["./first"]