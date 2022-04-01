# APP_NAME: Name of app ej. go-notes-project
# ACR: Name of ACR Service (extract from azure panel)
#
APP_NAME=$APP_NAME
ACR=$ACR_NAME
.PHONY: run
run:
	@go run cmd/api/main.go

.PHONY: build-pg
build-pg:
	@echo "Building docker image for postgres"
	@docker build --pull --rm -f "./dev/dockerfile-postgres-dev" -t pg-notes-container:latest "."

.PHONY: run-postgres
run-postgres:
	@echo "Running postgres container"
	@docker run -it -p 5432:5432 pg-notes-container:latest

.PHONY: run-mongo
run-mongo:
	@echo "Running mongodb container"
	@docker run -it -p 27017:27017 mongo

.PHONY: deploy
deploy:
	@echo "Deploy go app"
	@./deploy.sh $(APP_NAME) $(ACR)