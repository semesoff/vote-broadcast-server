package models

// Config models the configuration for the gateway service

type Route struct {
	Path   string `yaml:"path"`
	Method string `yaml:"method"`
}

type ServiceConfig struct {
	URL    string  `yaml:"url"`
	Routes []Route `yaml:"routes"`
}

type Config struct {
	Port         int                      `yaml:"port"`
	Services     map[string]ServiceConfig `yaml:"services"`
	JWTSecretKey JWTSecretKey             // from .env
}

type JWTSecretKey []byte
