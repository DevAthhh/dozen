package http

import (
	"context"
	"net"
	"net/http"

	handler "github.com/DevAthhh/DoZen/internal/http/handlers"
	"github.com/DevAthhh/DoZen/internal/repository"
	"github.com/DevAthhh/DoZen/pkg/lib/config"
	"go.uber.org/zap"
)

type repo interface {
	NewTaskRepo() *repository.TaskRepository
	NewUserRepo() *repository.UserRepository
	NewGroupRepo() *repository.GroupRepository
}

type server struct {
	httpServer *http.Server
}

func (s *server) Start() error {
	return s.httpServer.ListenAndServe()
}

func (s *server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func NewServer(logger *zap.Logger, repo repo) *server {
	addr := net.JoinHostPort(config.Cfg.HTTPServer.Host, config.Cfg.HTTPServer.Port)
	return &server{
		httpServer: &http.Server{
			Addr:         addr,
			Handler:      handler.Route(logger, repo),
			ReadTimeout:  config.Cfg.HTTPServer.RWTimeout,
			WriteTimeout: config.Cfg.HTTPServer.RWTimeout,
			IdleTimeout:  config.Cfg.HTTPServer.IdleTimeout,
		},
	}
}
