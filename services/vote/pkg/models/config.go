package models

type Config struct {
	GRPCServer          GRPCServer          `yaml:"GRPCServer"`
	Database            Database            `yaml:"database"`
	WebSocketGRPCServer WebSocketGRPCServer `yaml:"WebSocketGRPCServer"`
}

type GRPCServer struct {
	Port    string `yaml:"port"`
	Network string `yaml:"network"`
}

type WebSocketGRPCServer struct {
	Port    string `yaml:"port"`
	Network string `yaml:"network"`
	Url     string `yaml:"url"`
}

type Database struct {
	Host     string
	User     string
	Port     string
	Password string
	Db       string
	Driver   string
}
