# Makefile

.PHONY: serve

ENV ?= local

serve:
	@if [ "$(ENV)" = "development" ]; then \
		echo "Running in development mode"; \
	elif [ "$(ENV)" = "local" ]; then \
		echo "Running in local mode"; \
	else \
		echo "Running in $(ENV) mode"; \
	fi
	go run ./main.go $(ENV)

.PHONY: migration

migration:
	@read -p "Enter database type (mysql/mongodb) [default: mysql]: " dbtype; \
	if [ -z "$$dbtype" ]; then dbtype="mysql"; fi; \
	if [ "$$dbtype" != "mysql" ] && [ "$$dbtype" != "mongodb" ]; then \
		echo "Invalid database type. Please enter 'mysql' or 'mongodb'."; \
		exit 1; \
	fi; \
	read -p "Enter migration name: " name; \
	timestamp=$$(date +%Y%m%d%H%M%S); \
	filename="database/migrations/$$timestamp_$$dbtype_$$name.sql"; \
	mkdir -p database/migrations; \
	touch $$filename; \
	echo "Created migration file: $$filename"

.PHONY: prune

prune:
	@echo "Pruning processes using ports from .env.$(ENV)"
	@PORT=$$(grep -E '^PORT=' .env.$(ENV) | cut -d '=' -f2); \
	if [ -z "$$PORT" ]; then \
		echo "No PORT found in .env.$(ENV)"; \
	else \
		echo "Killing process on port $$PORT"; \
		fuser -cfu $$PORT/tcp || echo "No process found on port $$PORT"; \
	fi

.PHONY: restart
restart: prune serve