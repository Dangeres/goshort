build:
	go build -o bin/main cmd/server/main.go

run-all:
	docker-compose up --force-recreate --build 

lint:
	golangci-lint run