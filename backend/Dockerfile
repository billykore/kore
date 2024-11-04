FROM golang:1.22-alpine AS builder
WORKDIR /kore
COPY . .
ENV GOCACHE=/root/.cache/go-build
RUN go install github.com/google/wire/cmd/wire@latest && \
    go mod download && \
    wire ./cmd
RUN --mount=type=cache,target="/root/.cache/go-build" go build -order ./cmd/app ./cmd

FROM alpine:3.19
COPY --from=builder /kore/cmd/app /
EXPOSE 8000
ENTRYPOINT ["./app"]
