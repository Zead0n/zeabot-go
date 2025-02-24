FROM golang:1.22.2-alpine AS build
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go build -o /zeabot

FROM alpine:latest
WORKDIR /bot
COPY --from=build /zeabot .
ENTRYPOINT [ "./zeabot" ]
