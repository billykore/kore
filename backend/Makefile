# Examples:
# make build-image kind=gateway
# make build-image kind=service name=auth
.PHONY: build-image
build-image:
	@./scripts/build.sh $(kind) $(name)

# Examples:
# make deploy kind=gateway
# make deploy kind=service name=auth
# make deploy kind=database name=postgres
.PHONY: deploy
deploy:
	@./scripts/deployment.sh $(kind) $(name)

# Examples:
# make generate-wire service=product
.PHONY: generate-wire
generate-wire:
	@./scripts/wire.sh $(service)

# Examples:
# make generate-proto proto=auth.proto
.PHONY: generate-proto
generate-proto:
	@./scripts/grpc.sh $(proto)

# Examples:
# make generate
.PHONY: generate
generate: \
	generate-proto \
	generate-wire

# Examples:
# make run service=auth
.PHONY: run
run:
	@echo "Run app..."
	@go run ./services/$(service)/cmd

# Examples:
# make install-cli
.PHONY: install-cli
install-cli:
	@echo "Install monorepo CLI..."
	@go install ./korecli

# Examples:
# make gazelle-update-repos
.PHONY: gazelle-update-repos
gazelle-update-repos:
	@bazel run //:gazelle

# Examples:
# make bazel-run service=auth
.PHONY: bazel-run
bazel-run:
	@echo "Run $(service) service with bazel..."
	@bazel run //services/$(service)/cmd

# Examples:
# make create-migration name=create_table_products
.PHONY: create-migration
create-migration:
	@echo "Create migration..."
	@migrate create -ext sql -dir pkg/db/migrations -seq $(name)

# Examples:
# make migrate-up db=auth user=db_user password=db_password step=1
.PHONY: migrate-up
migrate-up:
	@echo "Migrate up..."
	@migrate -path pkg/db/migrations/$(db) -database 'postgres://$(user):$(password)@localhost:5432/$(db)?sslmode=disable' -verbose up $(step)

# Examples:
# make migrate-down db=auth user=db_user password=db_password step=1
.PHONY: migrate-down
migrate-down:
	@echo "Migrate down..."
	@migrate -path pkg/db/migrations/$(db) -database 'postgres://$(user):$(password)@localhost:5432/$(db)?sslmode=disable' -verbose down $(step)
