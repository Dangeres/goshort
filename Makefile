export REDIS_PASSWORD=12345
export REDIS_URL=redis://default:${REDIS_PASSWORD}@redis:6379/0?protocol=3

build:
	go build -o bin/main cmd/server/main.go

run-all:
	docker-compose up --force-recreate --build 

lint:
	golangci-lint run

gen-mocks:
	go generate ./internal/...

tests:
	go test ./internal/...

coverage:
	go test ./internal/... -coverprofile cover.out && go tool cover -func cover.out > coverage.out && \
	rm -f cover.out

coverage-html:
	go test -count=1 ./internal/... -coverprofile coverage.out && go tool cover -html=coverage.out && \
	rm coverage.out