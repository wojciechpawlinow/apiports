run:
	docker-compose build --no-cache
	docker-compose up -d && docker logs -f app_apiports

lint:
	go mod tidy
	go vet ./...
	go fmt ./...
	revive -config defaults.toml ./internal/...


test:
	go test -race ./...