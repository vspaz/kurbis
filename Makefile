BINARY_NAME=kurbis

all: build
build:
	go build -o $(BINARY_NAME) main.go

.PHONY: test
test:
	go test -race -v

.PHONY: clean
clean:
	rm -f $(BINARY_NAME)

.PHONY: style-fix
style-fix:
	gofmt -w .

.PHONE: lint
lint:
	golangci-lint run

.PHONY: upgrade
upgrade:
	go mod tidy
	go get -u all ./...
