package config

type AuthConfig struct {
	Enabled bool `yaml:"enabled"`
}

func NewAuthConfig(uri string) AuthConfig {
	svcConfig := NewServiceConfig(uri)
	return svcConfig.AuthConfig
}
