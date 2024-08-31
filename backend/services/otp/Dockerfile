FROM golang:1.22-alpine AS build
WORKDIR /app
COPY . .
RUN go install github.com/google/wire/cmd/wire@latest && \
    go mod download && \
    wire ./services/otp/... && \
    go build -o ./out/app ./services/otp/cmd

FROM alpine:3.19
COPY --from=build /app/out/app /
EXPOSE 8000
ENTRYPOINT ["./app"]
