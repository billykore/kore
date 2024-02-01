# make build-image service=todo
.PHONY: build-image
build-image:
	@./scripts/build.sh $(service)

# make deploy service=todo
.PHONY: deploy
deploy:
	@./scripts/deployment.sh $(service)

# make generate-wire
.PHONY: generate-wire
generate-wire:
	@./scripts/wire.sh

# make generate-grpc
.PHONY: generate-grpc
generate-grpc:
	@./scripts/grpc.sh

# make generate
.PHONY: generate
generate: \
	generate-grpc \
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
