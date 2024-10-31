package config

import (
	"flag"
	"os"

	"sync"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/rwx03/Pastebin/backend/pkg/logger"
)

type Config struct {
	Database Database `yaml:"database"`
	Server   Server   `yaml:"server"`
}

type Server struct {
	Port int `yaml:"port"`
}

type Database struct {
	Host         string `yaml:"host" env-default:"localhost"`
	Port         string `yaml:"port" env-default:"5432"`
	Username     string `yaml:"username" env-default:"postgres"`
	Password     string `yaml:"password" env-default:"123123"`
	DatabaseName string `yaml:"database_name" env-default:"attractive"`
}

var (
	once     sync.Once
	instance *Config
)

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}

		path := fetchConfigPath()
		if path == "" {
			logger.Log.Panicf("config path is empty")
		}

		if _, err := os.Stat(path); os.IsNotExist(err) {
			logger.Log.Panicf("config file does not exist: %s", path)
		}

		if err := cleanenv.ReadConfig(path, instance); err != nil {
			logger.Log.Fatalf("failed to read config, error: %s", err.Error())
		}
	})

	return instance
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
