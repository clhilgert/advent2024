BINARY_NAME=aoc
TARGET_DIR=./target
DEFAULT_GOAL := run

$(TARGET_DIR):
	@mkdir -p $(TARGET_DIR)

build: $(TARGET_DIR)
	GOOS=$(shell go env GOOS) GOARCH=$(shell go env GOARCH) go build -o $(TARGET_DIR)/$(BINARY_NAME) cmd/main/main.go

run: build
	./$(TARGET_DIR)/$(BINARY_NAME)

clean:
	rm -rf $(TARGET_DIR)
