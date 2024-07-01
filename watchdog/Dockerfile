FROM golang:1.20-alpine3.18 AS builder

# go_proxy
ARG go_proxy
ENV GOPROXY ${go_proxy}

# Workdir
WORKDIR /opt/monitor

# Copy file
COPY . .
# Build
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-w -s' -gcflags '-N -l' -o watchdog main.go

# Run
FROM alpine:3.18 AS runner
COPY --from=builder /opt/monitor/watchdog ./watchdog
ENTRYPOINT ["./watchdog"]

# docker build cesslab/watchdog:latest .
