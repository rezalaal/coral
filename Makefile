MIGRATIONS_PATH=$(shell pwd)/migrations
DB_URL=postgres://cafeuser:cafepass123@db:5432/cafedb?sslmode=disable
NETWORK_NAME=coral_cafe-net

migrate-up:
	docker run --rm -v $(MIGRATIONS_PATH):/migrations \
		--network $(NETWORK_NAME) \
		migrate/migrate \
		-path=/migrations \
		-database "$(DB_URL)" \
		up

migrate-down:
	docker run --rm -v $(PWD)/migrations:/migrations \
	  --network $(NETWORK_NAME) \
	  $(MIGRATE_IMAGE) -path=/migrations -database "$(DB_URL)" down 1

migrate-create:
	docker run --rm -v $(PWD)/migrations:/migrations \
	  $(MIGRATE_IMAGE) create -ext sql -dir /migrations -seq $(name)
