package tpl

func IRepositoryTemplate() []byte {
	return []byte(`package repository

type {{ .Repository }} interface {
	Get() error
}
`)
}
