package services

import (
	db2 "auth-service/internal/db"
	"auth-service/pkg/models"
)

type ServiceManager struct {
	Db db2.Database
}

func NewServiceManager(cfg models.Database) *ServiceManager {
	return &ServiceManager{
		Db: db2.NewDatabaseManager(cfg),
	}
}
