package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"vote-broadcast-server/services/vote/pkg/models"
)

type DatabaseManager struct {
	db *sql.DB
}

type Database interface {
	GetVotes(pollId int) (models.Votes, error)
	CreateVote(userVote models.UserVote) error
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
