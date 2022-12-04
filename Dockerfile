FROM golang:1.19.3

# simple - without multi-stage, alpine image and proper compilation
# just to show I worked on that. That is not being used for running

WORKDIR /app
COPY . .
RUN go mod download && go mod verify

ENTRYPOINT ["go", "run", "cmd/apiports/main.go"]