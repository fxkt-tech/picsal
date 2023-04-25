# *** params ***

APP_RELATIVE_PATH = $(shell a=`basename $$PWD` && cd .. && b=`basename $$PWD` && echo $$b/$$a)
INTERNAL_PROTO_FILES = $(shell find internal -name *.proto)
GEN_PROTO_FILES = $(shell find api -name "*.go")
PROTO_FILE_DIR = proto/api

COLOR = "\e[1;36m%s\e[0m\n"

# *** fast key ***

.PHONY: clean
clean: clean-api


# *** build ***

.PHONY: build
build:
	@printf $(COLOR) "Build: [$(APP_RELATIVE_PATH)]"
	@mkdir -p bin
	@go build -ldflags "-X main.Name=$(APP_NAME) -X main.Version=$(APP_VERSION)" -o ./bin/ ./...

# *** check ***

lint:
	@printf $(COLOR) "Lint: golangci-lint"
	@golangci-lint run

.PHONY: buf-lint
buf-lint:
	@printf $(COLOR) "Lint: buf-lint"
	@buf lint

# *** code gen ***

.PHONY: wire
wire:
	@printf $(COLOR) "Codegen: [$(APP_RELATIVE_PATH)] by wire"
	@cd cmd/server && wire

.PHONY: ent
ent:
	@printf $(COLOR) "Codegen: [$(APP_RELATIVE_PATH)] by ent"
	@cd internal/data && ent generate ./ent/schema

.PHONY: api
api: clean-api
	@printf $(COLOR) "Codegen: api"
	@buf generate --template buf.gen.yaml

.PHONY: openapi
openapi: 
	@printf $(COLOR) "Codegen: openapi"
	@buf generate --template buf.openapi.gen.yaml --path proto/api/cms/service

.PHONY: config
config:
	@printf $(COLOR) "Codegen: [$(APP_RELATIVE_PATH)] for config"
	@protoc --proto_path=. \
           --proto_path=../../../third_party \
           --go_out=paths=source_relative:. \
           $(INTERNAL_PROTO_FILES)

.PHONY: baseconfig
baseconfig:
	@printf $(COLOR) "Codegen: baseconfig"
	@cd third_party/yilan && protoc --proto_path=. \
		--go_out=paths=source_relative:. \
		./config/v1/*.proto

# *** tools ***

.PHONY: tools
tools:
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1
	go install entgo.io/ent/cmd/ent@v0.11.0
	go install github.com/bufbuild/buf/cmd/buf@v1.13.1
	cd pkg/kratos/cmd/protoc-gen-go-ylres && go install

# *** clean ***

.PHONY: clean-api
clean-api:
	@printf $(COLOR) "Clean: api gen code"
	@rm -rf api

.PHONY: test
test:
	go test -v ./... -cover