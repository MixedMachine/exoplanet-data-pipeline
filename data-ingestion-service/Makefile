.PHONY: pre build run test prod clean

VERSION=v1

pre:
	@cd src && go mod tidy
	@cd src && go fmt ./...

build.win: pre
	@cd src && GOOS=windows go build -o ../bin/data-ingestion-service.exe ./cmd/$(VERSION)/main.go

build.lin: pre
	@cd src && GOOS=linux go build -o ../bin/data-ingestion-service ./cmd/$(VERSION)/main.go

run: pre
	@cd src && go run ./cmd/$(VERSION)/main.go

test: pre
	@cd src && go test ./tests/...

prod:
	@docker build -t data-ingestion-service:$(VERSION) .