package services

import (
	"vote-service/internal/db"
	"vote-service/pkg/models"
)

type ServiceManager struct {
	Db db.Database
}

func NewServiceManager(cfg models.Database) *ServiceManager {
	return &ServiceManager{
		Db: db.NewDatabaseManager(cfg),
	}
}
