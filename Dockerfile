############################
# Build executable binary
############################
FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/lb

COPY . .
COPY /examples/example.yaml /examples/example.yaml

RUN go get -d -v
RUN ls -la

RUN go build -o /go/bin/lb

############################
# Build a small image from scratch
############################
FROM scratch

COPY --from=builder /go/bin/lb /go/bin/lb
COPY --from=builder /examples/example.yaml /examples/example.yaml

ENTRYPOINT ["/go/bin/lb"]