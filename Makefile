MIGRATIONS_DIR=migrations
CMD_DIR=cmd/v9

GO_EXEC=go

all: migrate run

migrate:
	@echo "Performing database migrations..."

# Target to run the Go application
run:
	cd $(CMD_DIR) && $(GO_EXEC) run .

# Phony targets
.PHONY: all migrate run
