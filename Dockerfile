FROM golang:1.15-alpine as builder
RUN apk --no-cache add git dep ca-certificates

ENV GOBIN=$GOPATH/bin
ENV GO111MODULE="on"

WORKDIR $GOPATH/src/github.com/diegosepusoto/nasa-graph-ql

COPY go.mod .
RUN go mod download
COPY src/ src/
COPY Makefile .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o $GOBIN/main src/main.go

FROM scratch as runner

ARG NASA_API_HOST
ARG NASA_API_KEY

ENV NASA_API_HOST=$NASA_API_HOST
ENV NASA_API_KEY=$NASA_API_KEY

EXPOSE 8080

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin /src/

ENTRYPOINT ["/src/main"]
