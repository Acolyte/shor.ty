### DOCKER ###

.PHONY: login
login:
	@docker login

.PHONY: pull
pull:
	@docker-compose pull $(PULL_ARGS)

.PHONY: up
up:
	@docker-compose up -d --remove-orphans --build

.PHONY: down
down:
	@docker-compose down --remove-orphans

.PHONY: logs
logs:
	@docker-compose logs

.PHONY: restart
restart:
	@docker-compose restart

.PHONY: build
build:
	@docker-compose build


### PLAIN BUILD ###

.PHONY: plain-swagger
plain-swagger:
	@go get -u github.com/swaggo/swag/cmd/swag && $(GOPATH)/bin/swag init -g ./internal/app/shorty/api/routers.go -o ./api --parseDependency --parseInternal --parseVendor

.PHONY: plain-assemble
plain-assemble:
	@go mod download && go build -ldflags "-s -w" -o ./bin/shorty ./cmd/shorty/main.go

.PHONY: plain-build
plain-build: plain-swagger plain-assemble

.PHONY: plain-run
plain-run:
	@cp ./.env ./bin/.env && ./bin/shorty

.PHONY: migrate-up
migrate-up:
	@docker-compose exec shorty bin/migrate -database "postgres://db:db@shorty-db:5432/db?sslmode=disable" -path /migration up

.PHONY: migrate-down
migrate-down:
	@docker-compose exec shorty bin/migrate -database "postgres://db:db@shorty-db:5432/db?sslmode=disable" -path /migration down

.PHONY: migrate
migrate: migrate-down migrate-up