run:
	docker run -d -p 27017:27017 mongo:latest
	sleep 5
	go run cmd/apiports/main.go # had to run go from host as I did not managed to make them working together as containers
								# although I tried with the same network and but Mongo were not reachable anyway

format:
	go mod tidy
	go vet ./...
	go fmt ./...
	revive -config defaults.toml ./internal/...


test:
	go test -race ./...