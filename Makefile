-include .env
export

CMD_DIR=cmd
APP=culinary_api

.PHONY: run
run:
	go run cmd/app/main.go

.PHONY: new-migration
new-migration:
	migrate create -ext sql -dir ./migrations/ -seq new_migration

.PHONY: build-linux
build-linux:
	CGO_ENABLED=0 GOARCH="amd64" GOOS=linux go build -ldflags="-s -w" -o ./bin/${APP} ${CMD_DIR}/app/main.go

.PHONY: swag_init
swag_init:
	swag init --parseDependency --dir ./internal/delivery/http/ -g router.go -o ./internal/delivery/http/docs


.PHONY: migrate-up
migrate-up:
	migrate -source file://migrations -database postgresql://sam:@localhost:5432/recipe_task?sslmode=disable up

