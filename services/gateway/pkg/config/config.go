package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"vote-broadcast-server/services/gateway/pkg/models"
)

type ConfigProvider interface {
	GetConfig() *models.Config
}

type ConfigManager struct {
	config *models.Config
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
	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			log.Fatalln("Error closing config: ", err)
		}
	}(file)

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
