package services

import (
	"poll-service/internal/db"
	"poll-service/pkg/models"
)

type ServiceManager struct {
	Db db.Database
}

func NewServiceManager(cfg models.Database) *ServiceManager {
	return &ServiceManager{
		Db: db.NewDatabaseManager(cfg),
	}
}
