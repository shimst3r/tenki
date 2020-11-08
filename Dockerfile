FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/shimst3r/tenki/
COPY . .
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/tenki
RUN chmod +x /go/bin/tenki

FROM scratch
COPY --from=builder /go/bin/tenki /go/bin/tenki
ENTRYPOINT ["/go/bin/tenki"]
CMD ["--help"]