package tpl

import "fmt"

func EntityTemplate() []byte {
	pkg := "package entity"
	req := "\ntype {{ .StructName }}Request struct {\n" +
		"\tName string `query:\"name\"`\n" +
		"}"
	res := "\ntype {{ .StructName }}Response struct {\n" +
		"\tMessage string `json:\"message\"`\n" +
		"}"
	s := fmt.Sprintf("%s\n%s\n%s", pkg, req, res)
	return []byte(s)
}
