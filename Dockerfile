FROM golang:latest AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=1

#Maintainer info
LABEL maintainer="@smf8"

#change it with your proxy server
#ARG http_proxy=http://192.241.180.51:443
#ARG https_proxy=http://192.241.180.51:443

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main .

#this step is for CGO libraries
RUN ldd main | tr -s '[:blank:]' '\n' | grep '^/' | \
    xargs -I % sh -c 'mkdir -p $(dirname ./%); cp % ./%;'
RUN mkdir -p lib64 && cp /lib64/ld-linux-x86-64.so.2 lib64/

#Second stage of build
FROM alpine
RUN apk update && apk --no-cache add ca-certificates

COPY --from=builder /build ./

ENTRYPOINT ["./main"]

