package models

type Config struct {
	GRPCServer      `yaml:"grpc_server"`
	WebSocketServer `yaml:"websocket_server"`
}

type GRPCServer struct {
	Port    string `yaml:"port"`
	Network string `yaml:"network"`
}

type WebSocketServer struct {
	Port    string `yaml:"port"`
	Network string `yaml:"network"`
}
