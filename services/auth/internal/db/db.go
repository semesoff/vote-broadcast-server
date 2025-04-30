package db

import (
	"auth-service/pkg/models"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type DatabaseManager struct {
	db *sql.DB
}

type Database interface {
	AddUser(user models.User) (models.UserWithID, error)
	GetUser(user models.User) (models.UserWithID, bool, error)
	GetUserWithPassword(user models.User) (models.UserWithPassword, bool, error)
}

func NewDatabaseManager(cfg models.Database) *DatabaseManager {
	db := &DatabaseManager{}
	db.Init(cfg)
	return db
}

func (d *DatabaseManager) Init(cfg models.Database) {
	var err error
	d.db, err = sql.Open(cfg.Driver, fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Db))
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
		return
	}
	if err = d.db.Ping(); err != nil {
		log.Fatalf("failed to ping database: %v", err)
		return
	}
	log.Println("database is initialized")
}
