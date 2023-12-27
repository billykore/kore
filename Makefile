# make build-image
.PHONY: build-image
build-image:
	./scripts/build.sh

# make deploy
.PHONY: deploy
deploy:
	./scripts/deployment.sh

# make generate-wire
.PHONY: generate-wire
generate-wire:
	./scripts/wire.sh

# make generate-grpc
.PHONY: generate-grpc
generate-grpc:
	./scripts/grpc.sh

# make generate
.PHONY: generate
generate: \
	generate-wire \
	generate-grpc
