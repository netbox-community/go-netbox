# Helpers
IS_DARWIN := $(filter Darwin,$(shell uname -s))

define set_env
	sed $(if $(IS_DARWIN),-i "",-i) -e "s/^#*\($(1)=\).*/$(if $(2),,#)\1$(2)/" .env
endef



# Environment recipes
.PHONY: default
default: init up

.PHONY: init
init:
	test -f .env || cp .env.example .env
	$(call set_env,USER_ID,$(shell id -u))

.PHONY: up
up:
	DOCKER_BUILDKIT=1 docker compose up -d --build

.PHONY: down
down:
	docker compose down

.PHONY: shell
shell:
	$(EXEC) zsh

# Project recipes
.PHONY: build
build:
	$(EXEC) ./scripts/set-versions.sh $(NETBOX_VERSION) $(NETBOX_DOCKER_VERSION)
	./scripts/fetch-spec.sh $$(cat api/netbox_version) $$(cat api/netbox_docker_version)
	$(EXEC) ./scripts/fix-spec.py
	./scripts/generate-code.sh
	$(EXEC) go mod tidy
	$(EXEC) goimports -w .

.PHONY: test
test:
	$(EXEC) go test -v ./...
