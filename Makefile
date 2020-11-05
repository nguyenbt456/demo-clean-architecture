.PHONY: build run test local-db

local-db:
	@docker-compose -f docker-compose/docker-compose_db.yml down
	@docker-compose -f docker-compose/docker-compose_db.yml up -d

	@docker cp ./database/db.sql pg_demo:./db.sql
	@docker exec -u postgres pg_demo psql demo postgres -f /db.sql

build:
	go build -o main .

run:
	go run main.go

test:
	go test ./...