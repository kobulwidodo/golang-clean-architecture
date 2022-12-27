.PHONY: swag-install
swag-install:
	@go install github.com/swaggo/swag/cmd/swag@v1.8.3

.PHONY: swaggo
swaggo:
	@/bin/rm -rf ./docs/swagger
	@`go env GOPATH`/bin/swag init -g ./src/cmd/main.go -o ./docs/swagger --pd --parseGoList -q

.PHONY: run
run: swaggo build
	@./build/app

.PHONY: run-tests
run-tests:
	@go test -v -failfast `go list ./... | grep -i 'business'` -cover

.PHONY: build
build:
	@go build -o ./build/app ./src/cmd