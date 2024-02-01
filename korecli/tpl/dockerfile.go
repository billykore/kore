package tpl

func DockerfileTemplate() []byte {
	return []byte(`FROM golang:1.21-alpine AS build
`)
}
