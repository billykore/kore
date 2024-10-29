package config

// firestore config.
type firestore struct {
	ProjectId       string `envconfig:"PROJECT_ID"`
	SDKFile         string `envconfig:"FIRESTORE_SDK"`
	TodoCollections string `envconfig:"TODO_COLLECTIONS"`
}
