.PHONY: backend

backend:
	@echo "Building backend..."
	@docker compose -f ./build/docker-compose.dev.yml up -d

backend.destroy:
	@echo "Breaking down backend..."
	@docker compose -f ./build/docker-compose.dev.yml down