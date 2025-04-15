package loadlogger

import (
	"log"

	"github.com/DevAthhh/DoZen/pkg/lib/config"
	"go.uber.org/zap"
)

func NewLogger() *zap.Logger {
	enviroment := config.Cfg.Env
	var logger *zap.Logger

	var err error
	switch enviroment {
	case config.Development:
		logger, err = zap.NewDevelopment()
	case config.Production:
		logger, err = zap.NewProduction()
	}

	if err != nil {
		log.Fatalf("error with initializing logger: %v", err)
	}

	return logger
}
