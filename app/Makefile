.PHONY = build test

LOG_DIR=./logs
CHECK_DIR=go list ./... | grep -v /cmd/utilits
SQL_DIR=./scripts
MICROSERVICE_DIR=$(PWD)/internal/microservices

stop-redis:
	systemctl stop redis

stop-postgres:
	systemctl stop postgresql

run-postgres-redis:
	systemctl start redis
	systemctl start postgresql


build:
	mkdir -p ./getme-secrt
	go build -o server.out -v ./cmd/server

build-docker: build-docker-server build-docker-pg

run-init:
	#sudo chown -R 5050:5050 ./pgadmin
	mkdir -p $(LOG_DIR)

run: run-init # запустить http сервер
	docker-compose -f docker-compose.yml --env-file ./configs/run-http.env up --build --no-deps

run-with-build: build-docker-server build-docker-pg run

stop:  # остановить сервер
	docker-compose stop

build-docker-server: # запуск обычного http servera
	docker build --no-cache --network host -f ./docker/builder.Dockerfile . --tag getme
build-docker-pg: # сборка образа базы
	docker build --no-cache --network host -f ./docker/postgresql.Dockerfile . --tag pg-14

open-last-log:
	cat $(LOG_DIR)/`ls -t $(LOG_DIR) | head -1 `

clear-logs:
	rm -rf $(LOG_DIR)/*.log

rm-docker:
	docker rm -vf $$(docker ps -a -q) || true

build-utils:
	go build -o utils.out -v ./cmd/utilits

parse-last-log: build-utils
	./utils.out -search-url=${search_url}

gen-mock:
	go generate -n $$(go list ./internal/...)

test:
	go test -v -race ./internal/...

DATABASE_URL:=$(shell cat ./configs/migrate.config | jq '.database_server')
DATABASE_URL_LOCAL:=$(shell cat ./configs/migrate.config | jq '.database_local')

migrate-up:
	migrate -source file://${SQL_DIR} -database ${DATABASE_URL} up

local-migrate-up:
	migrate -source file://${SQL_DIR} -database ${DATABASE_URL_LOCAL} up

migrate-down:
	migrate -source file://${SQL_DIR} -database ${DATABASE_URL} down

local-migrate-down:
	migrate -source file://${SQL_DIR} -database ${DATABASE_URL_LOCAL} down

migrate-up-one:
	migrate -source file://${SQL_DIR} -database ${DATABASE_URL} up 1

local-migrate-up-one:
	migrate -source file://${SQL_DIR} -database ${DATABASE_URL_LOCAL} up 1

migrate-down-one:
	migrate -source file://${SQL_DIR} -database ${DATABASE_URL} down 1

local-migrate-down-one:
	migrate -source file://${SQL_DIR} -database ${DATABASE_URL_LOCAL} down 1

migrate:
	migrate -source file://${SQL_DIR} -database ${DATABASE_URL} goto $(version)

local-migrate:
	migrate -source file://${SQL_DIR} -database ${DATABASE_URL_LOCAL} goto $(version)

force-migrate:
	migrate -source file://${SQL_DIR} -database ${DATABASE_URL} force $(version)

local-force-migrate:
	migrate -source file://${SQL_DIR} -database ${DATABASE_URL_LOCAL} force $(version)

generate:
	go generate -n $$(go list ./internal/...)