MIGRATIONS_PATH=$(shell pwd)/internal/db/migrations
DB_URL=postgres://cafeuser:cafepass123@db:5432/cafedb?sslmode=disable
PROJECT_NAME=$(shell basename $(shell pwd))
NETWORK_NAME=$(PROJECT_NAME)_cafe-net
MIGRATE_IMAGE=migrate/migrate

migrate-up:
	docker run --rm -v $(MIGRATIONS_PATH):/migrations \
		--network $(NETWORK_NAME) \
		$(MIGRATE_IMAGE) \
		-path=/migrations -database "$(DB_URL)" up

migrate-down:
	docker run --rm -v $(MIGRATIONS_PATH):/migrations \
		--network $(NETWORK_NAME) \
		$(MIGRATE_IMAGE) \
		-path=/migrations -database "$(DB_URL)" down 1

migrate-create:
	docker run --rm -v $(MIGRATIONS_PATH):/migrations \
		$(MIGRATE_IMAGE) \
		create -ext sql -dir /migrations -seq $(shell date +%y%m%d)

seed:
	sudo go run $(PWD)/cmd/seed/main.go

test:
	go test ./internal/user/repository/postgres
	go test ./internal/integration
	go test ./config
test-v:
	go test -v ./internal/user/repository/postgres
	go test -v ./internal/integration
	go test -v ./config
