package config

import (
	"flag"
	"os"
	"pastebin/pkg/logger"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	TelegramToken string   `yaml:"telegram-token"`
	Database      Database `yaml:"database"`
	Chats         Chats    `yaml:"chats"`
	Domains       Domains  `yaml:"domains"`
	Server        Server   `yaml:"server"`
}

type Server struct {
	Port int `yaml:"port"`
}

type Domains struct {
	Wallapop struct {
		ES string `yaml:"es"`
	} `yaml:"wallapop"`
}

type Database struct {
	Host         string `yaml:"host" env-default:"localhost"`
	Port         string `yaml:"port" env-default:"5432"`
	Username     string `yaml:"username" env-default:"postgres"`
	Password     string `yaml:"password" env-default:"123123"`
	DatabaseName string `yaml:"database_name" env-default:"attractive"`
}

type Chats struct {
	Applications int64 `yaml:"applications"`
	Logs         int64 `yaml:"logs"`
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
