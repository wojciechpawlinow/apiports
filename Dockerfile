FROM golang:1.19.4 AS builder

COPY . /go/src/github.com/wojciechpawlinow/apiports
WORKDIR /go/src/github.com/wojciechpawlinow/apiports

RUN go build -ldflags '-w -linkmode external -extldflags "-static"' -o bin/apiports ./cmd/apiports

FROM alpine

WORKDIR /app

COPY --from=builder --chown=app:app /go/src/github.com/wojciechpawlinow/apiports/bin/apiports /app/

RUN chmod +x /app/apiports

EXPOSE 8080

CMD ["/app/apiports"]