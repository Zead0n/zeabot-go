FROM golang:1.25-trixie AS build
WORKDIR /bot
RUN apt-get update && apt-get install -y pkg-config libopus-dev libopusfile-dev
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o bin/zeabot cmd/main.go

FROM debian:trixie-slim
RUN apt-get update && apt-get install -y \
    libopus0 \
    libopusfile0 \
    ca-certificates
RUN update-ca-certificates --fresh
ENV SSL_CERT_DIR=/etc/ssl/certs
COPY --from=build /bot/bin/zeabot /zeabot
ENTRYPOINT ["/zeabot"]
