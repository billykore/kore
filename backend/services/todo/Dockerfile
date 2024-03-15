FROM golang:1.21-alpine AS build
WORKDIR /app
COPY . .
RUN go install github.com/google/wire/cmd/wire@latest && \
    go mod download && \
    wire ./services/todo/... && \
    go build -o ./out/app ./services/todo/cmd

FROM alpine:3.19
COPY --from=build /app/out/app /
COPY --from=build /app/firebase-sdk.json /
EXPOSE 8000 9000
ENTRYPOINT ["./app"]
