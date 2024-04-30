export REDIS_PASSWORD=12345
export REDIS_URL=redis://default:${REDIS_PASSWORD}@redis:6379/0?protocol=3

build:
	go build -o bin/main cmd/server/main.go

run-all:
	docker-compose up --force-recreate --build 

lint:
	golangci-lint run

gen-mocks:
	go generate ./internal/domain/handlers
