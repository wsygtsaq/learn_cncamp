FROM golang:1.17 AS build
WORKDIR /httpserver/
COPY . .
ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV GOPATH=/go
ENV GOPROXY=https://goproxy.cn,direct
RUN GOOS=linux go build -o httpserver main.go

FROM busybox
COPY --from=build /httpserver/httpserver /httpserver/httpserver
EXPOSE 8891
WORKDIR /httpserver/
ENTRYPOINT ["./httpserver"]