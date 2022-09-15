# build file
GOCMD=go
# Use -a flag to prevent code cache problems.
GOBUILD=$(GOCMD) build -ldflags -s -v -i

BIN_BINARY_NAME=super-view-server
install:
	go mod tidy #fix missing go.sum entry for module providing package
	$(GOBUILD) -o $(BIN_BINARY_NAME) cmd/main.go
	@echo "Build $(BIN_BINARY_NAME) successfully. You can run ./$(BIN_BINARY_NAME) now.If you can't see it soon,wait some seconds"

update:
	go env -w GOPRIVATE="gitlab.***.me/**"
	go mod tidy
	#go mod vendor