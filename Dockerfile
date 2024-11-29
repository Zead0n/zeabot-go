FROM golang:1.22.2-alpine AS build
WORKDIR /build
RUN mkdir /bot
COPY . .
RUN go mod download
RUN go build -o /bot/zeabot

FROM alpine:latest
WORKDIR /bot
COPY --from=build /bot/zeabot .
ENTRYPOINT [ "./zeabot" ]
