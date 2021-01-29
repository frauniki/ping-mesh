FROM golang:1.15.7-alpine AS builder

WORKDIR /go/src/app
COPY . .

RUN go build -o ./bin/agent cmd/agent/main.go

FROM alpine:latest

COPY --from=builder /go/src/app/bin/agent /usr/local/bin

ENTRYPOINT [ "/usr/local/bin/agent" ]
