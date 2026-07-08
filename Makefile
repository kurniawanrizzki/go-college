BINARY_NAME := api
BIN_DIR     := ./bin
SRC_DIR     := ./cmd/api
CMD_PATH    := $(SRC_DIR)/main.go
DOCS_DIR    := ./api/openapi

GO          := go
SWAG        := swag

.PHONY: help swagger build run tidy clean

## Show available targets
help:
	@echo "Targets:"
	@echo "  make swagger   Generate Swagger API docs into $(DOCS_DIR)"
	@echo "  make build     Build the application into $(BIN_DIR)/$(BINARY_NAME)"
	@echo "  make run       Build and run the application"
	@echo "  make tidy      Tidy go modules"
	@echo "  make clean     Remove build artifacts"

## Generate swagger documentation
swagger:
	@echo "Generating Swagger docs..."
	@($(SWAG) fmt -d $(SRC_DIR) 2>&1 | grep -v "warning: failed to get package name in dir") || true
	@($(SWAG) init -g $(CMD_PATH) -o $(DOCS_DIR) 2>&1 | grep -v "warning: failed to get package name in dir") || true
	@sed -i.bak '/LeftDelim/d;/RightDelim/d' $(DOCS_DIR)/docs.go 2>/dev/null || true
	@rm -f $(DOCS_DIR)/docs.go.bak 2>/dev/null || true
	@echo "Swagger docs generated"

## Build the application
build: swagger
	@echo "Building $(BIN_DIR)/$(BINARY_NAME)..."
	@$(GO) build -o $(BIN_DIR)/$(BINARY_NAME) $(SRC_DIR)

## Build and run
run: build
	@$(BIN_DIR)/$(BINARY_NAME)

## Tidy modules
tidy:
	@$(GO) mod tidy

## Clean build artifacts
clean:
	@rm -rf $(BIN_DIR)
