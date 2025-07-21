MIGRATIONS_PATH=$(shell pwd)/internal/db/migrations
DB_URL=postgres://cafeuser:cafepass123@db:5432/cafedb?sslmode=disable
NETWORK_NAME=go-version_cafe-net
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
		create -ext sql -dir /migrations -seq $(name)

seed:
	sudo go run $(PWD)/cmd/seed/main.go

test:
	go test ./internal/repository/postgres
	go test ./internal/integration
test-v:
	go test -v ./internal/repository/postgres
	go test -v ./internal/integration
