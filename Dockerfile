############################
# Build executable binary
############################
FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/mypackage/myapp/

COPY . .

RUN go get -d -v

RUN go build -o /go/bin/lb

############################
# Build a small image from scratch
############################
FROM scratch

COPY --from=builder /go/bin/lb /go/bin/lb

ENTRYPOINT ["/go/bin/lb"]