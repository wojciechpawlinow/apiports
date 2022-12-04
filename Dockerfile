FROM golang:1.19.3

# simple - without multi-stage, alpine image and proper compilation

WORKDIR /app
COPY . .
RUN go mod download && go mod verify

# EXPOSE 8000 in case we provide an HTTP server

ENTRYPOINT ["go", "run", "cmd/apiports/main.go"]