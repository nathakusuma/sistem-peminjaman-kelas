# Use the .env file to load environment variables
-include .env

POSTGRES_URL := postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)

MIGRATE_CMD=migrate -database "${POSTGRES_URL}" -path database/migration/

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
