package config

type Firestore struct {
	ProjectId       string `envconfig:"PROJECT_ID"`
	SDKFile         string `envconfig:"FIRESTORE_SDK"`
	TodoCollections string `envconfig:"TODO_COLLECTIONS"`
}
