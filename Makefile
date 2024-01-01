BINARY_NAME="ginit"

build:
	@go build -o bin/$(BINARY_NAME) cmd/ginit/main.go

run: build
	@./bin/$(BINARY_NAME)

clean:
	@rm -rf bin/$(BINARY_NAME)