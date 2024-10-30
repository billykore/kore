package config

// token config.
type token struct {
	Secret    string `envconfig:"TOKEN_SECRET"`
	HeaderKid string `envconfig:"TOKEN_HEADER_KEY_ID"`
}
