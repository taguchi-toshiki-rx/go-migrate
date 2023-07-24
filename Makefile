include .env

.PHONY: install
install:
	curl -sSf https://atlasgo.sh | sh

.PHONY: up
up:
	docker compose up -d

.PHONY: down
down:
	docker compose down

.PHONY: run
run:
	docker exec -it go-migrate-sample-api ash -c "air"
