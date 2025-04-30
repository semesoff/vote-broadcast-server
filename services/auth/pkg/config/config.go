package config

import (
	"auth-service/pkg/models"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
	"log"
	"os"
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

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln("Error loading .env file: ", err)
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

	cm.config.Database = models.Database{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Db:       os.Getenv("DB_NAME"),
		Driver:   os.Getenv("DB_DRIVER"),
	}

	cm.config.JWTSecretKey = models.JWTSecretKey(os.Getenv("JWT_SECRET_KEY"))

	log.Println("Config is initialized")
}

func (cm *ConfigManager) GetConfig() *models.Config {
	return cm.config
}
