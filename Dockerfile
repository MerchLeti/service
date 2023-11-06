FROM golang:1.21-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o app cmd/etushop/main.go

FROM alpine
RUN apk add --no-cache openssl1.1-compat bash && adduser -D -h /home/container container

USER container
ENV USER=container HOME=/home/container GIN_MODE=release
WORKDIR /home/container

COPY --from=builder /build/app /app

COPY ./entrypoint.sh /entrypoint.sh
CMD ["/bin/bash", "/entrypoint.sh"]
