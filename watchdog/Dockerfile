FROM golang:1.20-alpine3.18 AS builder

ARG go_proxy
ENV GOPROXY ${go_proxy}

WORKDIR /opt/monitor

COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-w -s' -gcflags '-N -l' -o watchdog main.go

FROM alpine:3.18 AS runner
COPY --from=builder /opt/monitor/watchdog ./watchdog
ENTRYPOINT ["./watchdog"]

# docker build cesslab/watchdog:latest .
