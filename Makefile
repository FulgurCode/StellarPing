build:
	@templ generate
	@go build -o ./bin/stellar-ping ./cmd/main.go

run: build
	@./bin/stellar-ping
