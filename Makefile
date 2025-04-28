BINARY_NAME=tdoro
DIST_DIR=dist

.PHONY: all clean build-linux build-macos-amd64 build-macos-arm64

all: clean build-linux build-macos-amd64 build-macos-arm64

build-linux:
	GOOS=linux GOARCH=amd64 go build -o $(DIST_DIR)/$(BINARY_NAME)-linux-amd64

build-macos-amd64:
	GOOS=darwin GOARCH=amd64 go build -o $(DIST_DIR)/$(BINARY_NAME)-darwin-amd64

build-macos-arm64:
	GOOS=darwin GOARCH=arm64 go build -o $(DIST_DIR)/$(BINARY_NAME)-darwin-arm64

clean:
	rm -rf $(DIST_DIR)
	mkdir -p $(DIST_DIR)
