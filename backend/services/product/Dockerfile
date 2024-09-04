FROM golang:1.22-alpine AS build
WORKDIR /app
COPY . .
ENV GOCACHE=/root/.cache/go-build
RUN go install github.com/google/wire/cmd/wire@latest && \
    go mod download && \
    wire ./services/product/...
RUN --mount=type=cache,target="/root/.cache/go-build" go build -o ./out/app ./services/product/cmd

FROM alpine:3.19
COPY --from=build /app/out/app /
EXPOSE 8000
ENTRYPOINT ["./app"]
