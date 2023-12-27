FROM golang:1.21-alpine AS build
WORKDIR /app
COPY . .
RUN apk add --no-cache protobuf && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 && \
    go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest && \
    go install github.com/google/wire/cmd/wire@latest && \
    go mod download && \
    protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        --grpc-gateway_out=. --grpc-gateway_opt=paths=source_relative \
        internal/grpc/v1/todo.proto && \
    wire ./... && \
    go build -o ./out/app ./cmd

FROM alpine:latest
COPY --from=build /app/out/app /
COPY --from=build /app/todo-list-app-firebase-sdk.json /
EXPOSE 8000
EXPOSE 9000
ENTRYPOINT ["./app"]
