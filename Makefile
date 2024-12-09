build:
	@go build -o bin/AuthSystem ./cmd/main.go

run: build
	@./bin/AuthSystem

