package app

import (
	"log"

	"github.com/DevAthhh/DoZen/internal/http"
	"github.com/DevAthhh/DoZen/internal/repository"
	"go.uber.org/zap"
)

type Repo interface {
	NewTaskRepo() *repository.TaskRepository
	NewUserRepo() *repository.UserRepository
	NewGroupRepo() *repository.GroupRepository
}

func Run(logger *zap.Logger, repo Repo) {
	srv := http.NewServer(logger, repo)
	logger.Debug("Server has been started")
	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
