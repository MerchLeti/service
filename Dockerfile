FROM golang:1.21-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o app cmd/etushop/main.go

FROM node:20-alpine AS frontend_builder

RUN mkdir -p /home/node/app/node_modules && chown -R node:node /home/node/app
WORKDIR /home/node/app
USER node
COPY --chown=node:node website/leti-merch .
RUN npm install
RUN npm run build

FROM alpine
RUN apk add --no-cache openssl1.1-compat bash && adduser -D -h /home/container container

USER container
ENV USER=container HOME=/home/container GIN_MODE=release
WORKDIR /home/container

COPY --from=builder /build/app /app
COPY --from=frontend_builder /home/node/app/dist /frontend

COPY ./entrypoint.sh /entrypoint.sh
CMD ["/bin/bash", "/entrypoint.sh"]
