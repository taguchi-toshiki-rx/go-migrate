include .env

MIGRATION_CONTAINER=go-migrate-sample-db-migration

.PHONY: up
up:
	docker compose up -d

.PHONY: down
down:
	docker compose down

.PHONY: run
run:
	docker exec -it go-migrate-sample-api ash -c "air"

.PHONY: migration-create
migration-create:
ifndef name
	@echo "Usage: make create name=create_users_table"
	@echo ""
	@exit 1
endif
	docker compose run --rm --no-deps --entrypoint migrate ${MIGRATION_CONTAINER} create -ext sql -seq -dir migrations $(name)

.PHONY: migration-up
migration-up:
	docker compose run --rm ${MIGRATION_CONTAINER}

.PHONY: migration-down
migration-down:
	docker compose run --rm ${MIGRATION_CONTAINER} down
