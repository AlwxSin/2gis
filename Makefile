.PHONY: test migrate

ci: test lint

run:
	go build && ./applicationDesignTest

test:
	go test ./... -p 1 -cover -coverprofile=coverage.out -coverpkg=./...

coverage:
	go tool cover -html=coverage.out

lint:
	golangci-lint run

# use with caution
# can delete nolint comments
# can mess with imports
lintfix:
	golangci-lint run --fix

fmt:
	gofumpt -l -w .
	gci write . --skip-generated -s standard -s default -s "prefix(applicationDesignTest)" -s blank -s dot
