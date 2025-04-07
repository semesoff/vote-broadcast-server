package models

type Config struct {
	GRPCServer   GRPCServer   `yaml:"GRPCServer"`
	Database     Database     `yaml:"database"`
	JWTSecretKey JWTSecretKey // form .env
}

type GRPCServer struct {
	Port    string `yaml:"port"`
	Network string `yaml:"network"`
}

type Database struct {
	Host     string
	User     string
	Port     string
	Password string
	Db       string
	Driver   string
}
