package services

import (
	"vote-broadcast-server/services/vote/internal/db"
	"vote-broadcast-server/services/vote/pkg/models"
)

type ServiceManager struct {
	Db db.Database
}

func NewServiceManager(cfg models.Database) *ServiceManager {
	return &ServiceManager{
		Db: db.NewDatabaseManager(cfg),
	}
}
