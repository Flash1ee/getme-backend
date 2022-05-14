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


build: build-swagger
	mkdir -p ./getme-secrt
	go build -o server.out -v ./cmd/server

build-docker: build-swagger build-docker-server build-docker-pg build-docker-sessions

run-init:
	#sudo chown -R 5050:5050 ./pgadmin
	mkdir -p $(LOG_DIR)

run: run-init # запустить http сервер
	docker-compose -f docker-compose.yml --env-file ./configs/run-http.env up --build --no-deps

run-with-build: build-docker run

stop:  # остановить сервер
	docker-compose stop

build-docker-server: # запуск обычного http servera
	docker build --network host -f ./docker/builder.Dockerfile . --tag getme-main
build-docker-pg: # сборка образа базы
	docker build  --network host -f ./docker/postgresql.Dockerfile . --tag pg-14

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

gen-proto-sessions:
	#protoc --proto_path=${MICROSERVICE_DIR}/auth/delivery/grpc/protobuf session.proto --go-grpc_out=plugins=grpc:${MICROSERVICE_DIR}/auth/delivery/grpc/protobuf
	protoc -I=${MICROSERVICE_DIR}/auth/delivery/grpc/protobuf --go_out=:${MICROSERVICE_DIR}/auth/delivery/grpc/protobuf --go_opt=paths=source_relative \
    --go-grpc_out=${MICROSERVICE_DIR}/auth/delivery/grpc/protobuf --go-grpc_opt=require_unimplemented_servers=false ${MICROSERVICE_DIR}/auth/delivery/grpc/protobuf/session.proto
	#protoc -I=${MICROSERVICE_DIR}/auth/delivery/grpc/protobuf --go_out=paths=source_relative:${MICROSERVICE_DIR}/auth/delivery/grpc/protobuf \
#	${MICROSERVICE_DIR}/auth/delivery/grpc/protobuf/session.proto  --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative

run-coverage:
	go test -covermode=atomic -coverprofile=cover ./...
	cat cover | fgrep -v "easyjson" | fgrep -v "mock" | fgrep -v "testing.go" | fgrep -v "docs" | fgrep -v ".pb.go" | fgrep -v "config" > cover2
	go tool cover -func=cover2

build-swagger:
	go run ./tools/swagger/prepare.go
	cp -r ./api/swagger.json ./swagger/swagger-ui/.
	cp -r ./api/swagger.json ./swagger/swagger-editor/.
	go list github.com/rakyll/statik || go get github.com/rakyll/statik
	statik -src=./swagger

build-editor-swagger:
	cp -r ./api/swagger.json ./swagger-editor/node_modules/swagger-editor-dist/.

## Microservices

build-sessions:
	go build -o sessions.out -v ./cmd/sessions

build-docker-sessions: # сборка образа сервиса авторизаций
	docker build  --network host -f ./docker/services/session-service.Dockerfile . --tag session-service