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
	go run cmd/app/main.go $(ENV)
