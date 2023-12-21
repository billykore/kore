FROM golang:1.21-alpine AS build
WORKDIR /app
COPY . .
RUN go mod download && \
    go install github.com/google/wire/cmd/wire@latest && \
    wire ./... && \
    go build -o ./out/app ./cmd

FROM alpine:latest
COPY --from=build /app/out/app /
COPY --from=build /app/todo-list-app-firebase-sdk.json /

EXPOSE 8080
ENTRYPOINT ["./app"]
