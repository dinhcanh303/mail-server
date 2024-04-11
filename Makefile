include .env
export

all: build test

sqlc: 
	sqlc generate
.PHONY: sqlc

test:
	go test -v main.go

clean:
	go clean

clean-database:
	docker volume rm $$(docker volume ls -q)

linter-golangci: ### check by golangci linter
	golangci-lint run
.PHONY: linter-golangci

createdb:
	docker exec -it mcs-postgres createdb --username=postgres --owner=root postgres

dropdb:
	docker exec -it mcs-postgres dropdb postgres

migrate-refresh: migrate-down migrate-up

migrate-up:
	migrate -path db/migrations -database "$(PG_URL)" -verbose up

migrate-down:
	migrate -path db/migrations -database "$(PG_URL)" -verbose down

.PHONY: createdb dropdb migrate-up migrate-down migrate-refresh

db_docs:
	dbdocs build docs/db.dbml

db_schema:
	dbml2sql --postgres -o docs/schema.sql docs/db.dbml

wire:
	cd internal/mail/app && wire && cd -
.PHONY: wire

proto:
	buf generate
	
proto-gen:
	rm -f proto/gen/*.go
	rm -f third_party/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=proto/gen --go_opt=paths=source_relative \
	--go-grpc_out=proto/gen --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=proto/gen --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=third_party/swagger --openapiv2_opt=allow_merge=true,merge_file_name=mail-server\
	proto/*.proto
	statik -src=./third_party/swagger -dest=./third_party
.PHONY: proto

docker: docker-stop docker-start
.PHONY: docker

docker-start:
	docker-compose up --build
.PHONY: docker-start

docker-stop:
	docker-compose down
.PHONY: docker-stop

docker-core: docker-core-stop docker-core-start

docker-core-start:
	docker-compose -f docker-compose-core.yaml up --build -d
.PHONY: docker-core-start

docker-core-stop:
	docker-compose -f docker-compose-core.yaml down
# --remove-orphans -v
.PHONY: docker-core-stop

docker-build:
	docker-compose down --remove-orphans -v
	docker-compose build
.PHONY: docker-build

run: run-mail run-proxy

run-proxy:
	cd cmd/proxy && go mod tidy && go mod download && \
	CGO_ENABLED=0 go run main.go &
.PHONY: run-proxy

run-mail:
	cd cmd/mail && go mod tidy && go mod download && \
	CGO_ENABLE=0 go run main.go &
.PHONY: run-mail



