# --------------------------------------------------------------------------------------------------
# API

.PHONY: help ## Print this help message
help:
	@echo "*** help ***"
	@awk '/^.PHONY/ { printf("%-10s\t", $$2); for(i=4; i<NF; i++) { printf("%s ", $$i); } printf("\n"); }' <Makefile

.PHONY: run ## Run the spirits API on port 8080
run: generate
	docker build -t spirits:dev .
	docker run --rm -it -p 8080:80 spirits:dev

GENERATED_DIRS=\
  pkg/generated/client \
  pkg/generated/server \
  script/generated/cli \
  doc/generated/adoc \
  doc/generated/html

.PHONY: generate ## Generate generated files
generate: $(GENERATED_DIRS)

# --------------------------------------------------------------------------------------------------
# internal

GENERATED_API=pkg/generated/api.json

OPENAPI_GENERATE_PREFIX=\
  docker run --rm -v $(shell pwd):/local openapitools/openapi-generator-cli generate \
    -v \
    -i /local/$< \
    -o /local/$@ \
    -g

$(GENERATED_API): cmd/genapi
	go run ./$< >$@

pkg/generated/client: $(GENERATED_API)
	$(OPENAPI_GENERATE_PREFIX) go

pkg/generated/server: $(GENERATED_API)
	$(OPENAPI_GENERATE_PREFIX) go-server

script/generated/cli: $(GENERATED_API)
	$(OPENAPI_GENERATE_PREFIX) bash

doc/generated/adoc: $(GENERATED_API)
	$(OPENAPI_GENERATE_PREFIX) asciidoc

doc/generated/html: $(GENERATED_API)
	$(OPENAPI_GENERATE_PREFIX) html
