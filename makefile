.PHONY: build
build:
	@go build -o ./build/app ./src/cmd

.PHONY: run
run:
	@go run ./src/cmd/main.go