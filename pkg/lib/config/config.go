package config

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
)

type httpServer struct {
	Port        string
	Host        string
	IdleTimeout time.Duration
	RWTimeout   time.Duration
}

type Config struct {
	Env        string
	APIVersion string
	HTTPServer httpServer
}

var Cfg *Config

func MustLoad() {
	path := os.Getenv("PATH_TO_CONFIG")
	if path == "" {
		log.Fatal("path to config is missing")
	}
	name := os.Getenv("CONFIG_NAME")
	if name == "" {
		name = "config_local.yml"
	}

	pathWithFile := filepath.Join(path, name)

	if _, err := os.Stat(pathWithFile); err != nil {
		log.Fatalf("the config file is missing in the specified path: %v", err.Error())
	}

	v := viper.New()
	v.SetConfigFile(pathWithFile)

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("error when reading the config file: %v", err.Error())
	}

	if v.GetString("server.port") == "" || v.GetString("server.host") == "" {
		log.Fatal("port and host cannot be empty")
	}

	Cfg = &Config{
		Env:        v.GetString("env"),
		APIVersion: v.GetString("api.version"),
		HTTPServer: httpServer{
			Port:        v.GetString("server.port"),
			Host:        v.GetString("server.host"),
			RWTimeout:   v.GetDuration("server.rw_timeout"),
			IdleTimeout: v.GetDuration("server.idle_timeout"),
		},
	}
}
