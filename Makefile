.PHONY: run
run:
	@go run cmd/api/main.go

.PHONY: build-pg
build-pg:
	@echo "Construyendo imagen de docker para postgres"
	@docker build --pull --rm -f "./dev/dockerfile-postgres-dev" -t pg-notes-container:latest "."

.PHONY: run-postgres
run-postgres:
	@echo "Corriendo contenedor de postgres"
	@docker run -it -p 5432:5432 pg-notes-container:latest

.PHONY: run-mongo
run-mongo:
	@echo "Corriendo contenedor de mongodb"
	@docker run -it -p 27017:27017 mongo