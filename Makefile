.PHONY: run
run:
	@go run ./src/cmd/main.go

.PHONY: swaggo
swaggo:
	@/bin/rm -rf ./docs/swagger
	@`go env GOPATH`/bin/swag init -g ./src/cmd/main.go -o ./docs/swagger --parseInternal


.PHONY: swag-install
swag-install:
	@go install github.com/swaggo/swag/cmd/swag@v1.6.7

.PHONE: run-app
run-app:
	@make swaggo
	@make run