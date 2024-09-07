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

# make generate-wire service=product
.PHONY: generate-wire
generate-wire:
	@./scripts/wire.sh $(service)

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

# make bazel-run service=todo
.PHONY: bazel-run
bazel-run:
	@echo "Run $(service) service with bazel..."
	@bazel run //services/$(service)/cmd

# make create-migration name=create_table_products
.PHONY: create-migration
create-migration:
	@echo "Create migration..."
	@migrate create -ext sql -dir pkg/db/migrations -seq $(name)

# make migrate-up dsn=postgres://user:password@localhost:5432/db?sslmode=disable step=1
.PHONY: migrate-up
migrate-up:
	@echo "Migrate up..."
	@migrate -path pkg/db/migrations -database $(dsn) -verbose up $(step)

# make migrate-down dsn=postgres://user:password@localhost:5432/db?sslmode=disable step=1
.PHONY: migrate-down
migrate-down:
	@echo "Migrate down..."
	@migrate -path pkg/db/migrations -database $(dsn) -verbose down $(step)
