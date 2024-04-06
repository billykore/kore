# make build-gateway-image
.PHONY: build-gateway-image
build-gateway-image:
	@./scripts/build.sh "gateway"

# make build-image service=todo
.PHONY: build-image
build-image:
	@./scripts/build.sh "service" $(service)

# make deploy-gateway
.PHONY: deploy-gateway
deploy-gateway:
	@./scripts/deployment.sh "gateway"

# make deploy service=todo
.PHONY: deploy
deploy:
	@./scripts/deployment.sh "service" $(service)

# make generate-wire
.PHONY: generate-wire
generate-wire:
	@./scripts/wire.sh

# make generate-proto proto=todo.proto
.PHONY: generate-proto
generate-proto:
	@./scripts/grpc.sh $(proto)

# make generate
.PHONY: generate
generate: \
	generate-proto \
	generate-wire

# make run service=todo
.PHONY: run
run:
	@echo "Run app..."
	@go run ./services/$(service)/cmd

# make install-cli
.PHONY: install-cli
install-cli:
	@echo "Install monorepo CLI..."
	@go install ./korecli

# make gazelle-update-repos
.PHONY: gazelle-update-repos
gazelle-update-repos:
	@bazel run //:gazelle

# make load-env
.PHONY: load-env
load-env:
	@./scripts/env.sh

# make bazel-run service=todo
.PHONY: bazel-run
bazel-run: load-env
	@echo "Run $(service) service with bazel..."
	@bazel run //services/$(service)/cmd
