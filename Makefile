run:
	go run ./cmd/main.go

build:
	go build -o ./bin/fl_struct ./cmd/main.go

run-linter:
	golangci-lint run ./cmd/main.go

fmt:
	gofumpt -l -w ./..