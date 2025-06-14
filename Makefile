# Use the .env file to load environment variables
-include .env

POSTGRES_URL := postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

MIGRATE_CMD=docker compose run --rm migrate -database "${POSTGRES_URL}" -path migration/
SEED_CMD=docker compose exec -T db psql -U $(DB_USER) -d $(DB_NAME) -W $(DB_PASS) < database/seeder/

# Targets for different migration commands
.PHONY: up down status version force

# Apply all migrations
migrate-up:
	$(MIGRATE_CMD) up

# Rollback the most recent migration
migrate-down:
	$(MIGRATE_CMD) down

# Show the status of migrations
migrate-status:
	$(MIGRATE_CMD) status

# Show the current version of the database
migrate-version:
	$(MIGRATE_CMD) version

# Force a specific version (replace <version> with the desired version)
migrate-force:
	$(MIGRATE_CMD) force $(version)

seed-up:
	$(SEED_CMD)up.sql

seed-down:
	$(SEED_CMD)down.sql
