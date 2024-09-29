export GO111MODULE=on
export GOBIN=$(CURDIR)/bin
export BUF_BIN=$(GOBIN)/buf
export GOOSE_BIN?=$(GOBIN)/goose
export LOCAL_BIN:=$(CURDIR)/bin
export MIGRATIONS_DIR=./migrations

bin-deps:
	$(info Installing binary dependencies...)
	go install github.com/yesmishgan/protoc-gen-bomboglot@v0.0.1
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	go install github.com/bufbuild/buf/cmd/buf@v1.31.0

generate:
	$(BUF_BIN) generate

GOLANGCI_BIN := $(LOCAL_BIN)/golangci-lint
GOLANGCI_TAG ?= 1.55.2

install-lint: export GOBIN := $(LOCAL_BIN)
install-lint: ## Установить golangci-lint в текущую директорию с исполняемыми файлами
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v$(GOLANGCI_TAG)

.lint: install-lint
	$(GOLANGCI_BIN) run \
		--sort-results \
		--max-issues-per-linter=1000 \
		--max-same-issues=1000 \
		./...

.PHONY: lint
lint: .lint

install-goose:
	go install github.com/pressly/goose/v3/cmd/goose@latest

.PHONY: migration
migration: install-goose
	$(GOOSE_BIN) -dir migrations postgres \
				"host=localhost \
				port=5432 \
				dbname=test \
				user=test \
				password=test \
				sslmode=disable" up

.PHONY: add-new-migration
add-new-migration:
	goose --dir=$(MIGRATIONS_DIR) create new_migration sql

