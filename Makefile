# --------------------------------------------------------------------------------------------------
# API

SHELL=bash -o pipefail

VERSION=0.0.1

GENERATED_API=pkg/api/generated/api.json

GENERATED_DIRS=\
  pkg/api/generated/client \
  pkg/api/generated/server \
  script/generated/cli \
  doc/generated/adoc \
  doc/generated/html

OPENAPI_GENERATE_PREFIX=\
  docker run --rm -v $(shell pwd):/local openapitools/openapi-generator-cli:latest-release generate \
    -i /local/$< \
    -o /local/$@ \
    -v \
    -g

OPENAPI_VALIDATE_PREFIX=\
  docker run --rm -v $(shell pwd):/local openapitools/openapi-generator-cli:latest-release validate \
    -i /local/$<

.PHONY: help ## Print this help message
help:
	@echo "*** help ***"
	@awk '/^.PHONY/ { printf("%-10s\t", $$2); for(i=4; i<=NF; i++) { printf("%s ", $$i); } printf("\n"); }' <Makefile

.PHONY: run ## Run the spirits API on port 8080
run: generate
	docker build -t spirits:dev .
	docker run --rm -it -p 8080:80 spirits:dev

.PHONY: generate ## Generate generated files
generate: $(GENERATED_DIRS)

.PHONY: validate ## Validate API spec
validate: $(GENERATED_API)
	$(OPENAPI_VALIDATE_PREFIX)

# --------------------------------------------------------------------------------------------------
# internal

$(GENERATED_API): cmd/genapi/main.go cmd/genapi/config.json cmd/genapi/api.json.tmpl
	go run ./$< $(VERSION) | jq . >$@

COMMON_GO_CONFIG_OPTIONS=enumClassPrefix=true,hideGenerationTimestamp=false,isGoSubmodule=true,packageVersion=$(VERSION),

pkg/api/generated/client: $(GENERATED_API)
	rm -rf $@
	$(OPENAPI_GENERATE_PREFIX) go -p $(COMMON_GO_CONFIG_OPTIONS),packageName=client

pkg/api/generated/server: $(GENERATED_API)
	rm -rf $@
	$(OPENAPI_GENERATE_PREFIX) go-server -p $(COMMON_GO_CONFIG_OPTIONS),onlyInterfaces=true,outputAsLibrary=true,packageName=server,sourceFolder=api
	goimports -w $@

script/generated/cli: $(GENERATED_API)
	rm -rf $@
	$(OPENAPI_GENERATE_PREFIX) bash

doc/generated/adoc: $(GENERATED_API)
	rm -rf $@
	$(OPENAPI_GENERATE_PREFIX) asciidoc

doc/generated/html: $(GENERATED_API)
	rm -rf $@
	$(OPENAPI_GENERATE_PREFIX) html2
