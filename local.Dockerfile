FROM golang:1.21-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o app cmd/etushop/main.go

FROM alpine

WORKDIR /

COPY --from=builder /build/app /app

CMD ["/app"]
