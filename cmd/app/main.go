package main

import (
	"github.com/DevAthhh/DoZen/internal/app"
	"github.com/DevAthhh/DoZen/internal/lib/database"
	"github.com/DevAthhh/DoZen/internal/repository"
	"github.com/DevAthhh/DoZen/pkg/lib/config"
	loadenv "github.com/DevAthhh/DoZen/pkg/lib/loadEnv"
	loadlogger "github.com/DevAthhh/DoZen/pkg/lib/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// TODO: сделать структуру реквестов и респонсов

var (
	logger *zap.Logger
	db     *gorm.DB

	repo *repository.Repository
)

func init() {
	loadenv.LoadEnv()
	config.MustLoad()
	logger = loadlogger.NewLogger()

	db = database.LoadDatabase()
	database.SyncDB(db)

	repo = repository.NewRepository(db)
}

func main() {
	defer logger.Sync()
	defer func() {
		if db, err := db.DB(); err == nil {
			db.Close()
		}
	}()

	app.Run(logger, repo)
}
