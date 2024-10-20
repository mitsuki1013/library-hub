DATABASE_URL := "mysql://root:root@tcp(db:3306)/library_hub"
MIGRATOR := docker compose exec migrator
MIGRATIONS_PATH := "/migrations"

reset-volumes:
	@docker compose down -v; \
	docker compose up -d --wait

migrate-up:
	$(MIGRATOR) migrate -database $(DATABASE_URL) -path $(MIGRATIONS_PATH) up

# make migrate-down STEPS=N
migrate-down:
	@if [ -z "$(STEPS)" ]; then \
		echo "Usage: make migrate-down steps=<number_of_steps>"; \
	else \
		$(MIGRATOR) migrate -database $(DATABASE_URL) -path $(MIGRATIONS_PATH) down $(STEPS); \
	fi

# make create-migration NAME=xxx
create-migration:
	$(MIGRATOR) migrate create -ext sql -dir $(MIGRATIONS_PATH) -seq "$(NAME)"