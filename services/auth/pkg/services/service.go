package services

import (
	db2 "vote-broadcast-server/services/auth/internal/db"
	"vote-broadcast-server/services/auth/pkg/models"
)

type ServiceManager struct {
	Db db2.Database
}

func NewServiceManager(cfg models.Database) *ServiceManager {
	return &ServiceManager{
		Db: db2.NewDatabaseManager(cfg),
	}
}
