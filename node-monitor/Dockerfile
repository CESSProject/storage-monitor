FROM golang:1.20-alpine3.18 AS builder

# go_proxy
ARG go_proxy
ENV GOPROXY ${go_proxy}

# Workdir
WORKDIR /opt/target

# Copy file
COPY ./server ./server
COPY ./dashboard/dist ./dist
# Build
RUN cd server && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-w -s' -gcflags '-N -l' -o node-monitor main.go

# Run
FROM alpine:3.18 AS runner
WORKDIR /opt/server
COPY --from=builder /opt/target/server/node-monitor ./node-monitor
COPY --from=builder /opt/target/dist ./dist
ENTRYPOINT ["./node-monitor"]