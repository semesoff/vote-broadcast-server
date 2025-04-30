package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"websocket-service/pkg/models"
)

type ConfigManager struct {
	config *models.Config
}

type ConfigProvider interface {
	GetConfig() *models.Config
}

func NewConfigManager() *ConfigManager {
	cfg := &ConfigManager{config: &models.Config{}}
	cfg.Init()
	return cfg
}

func (cm *ConfigManager) Init() {
	file, err := os.Open("pkg/config/config.yaml")
	if err != nil {
		log.Fatalln("Error opening config: ", err)
		return
	}

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(cm.config); err != nil {
		log.Fatalln("Error decoding config: ", err)
		return
	}

	log.Println("Config is initialized")
}

func (cm *ConfigManager) GetConfig() *models.Config {
	return cm.config
}
