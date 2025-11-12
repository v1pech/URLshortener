package config

import (
	"errors"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	StoragePath string `yaml:"storagePath" env-default:"../storage/sqlite/storage.db"`
	LogPath     string `yaml:"logPath" env-default:"../logs/logs.log"`
	LogLevel    string `yaml:"logLevel" env-default:"info"`
	Server      Server `yaml:"server"`
}

type Server struct {
	Adress      string `yaml:"address" env-default:"localhost:8080"`
	Timeout     int    `yaml:"timeout" env-default:"4"`
	IdleTimeout int    `yaml:"idleTimeout" env-default:"45"`
	User        string `yaml:"user" env-default:"admin"`
	Password    string `yaml:"password" env-required:"true" env:"URLSHORTENER_PASSWORD"`
}

func MustLoad() (*Config, error) {

	var config Config

	err := cleanenv.ReadConfig("../config/config.yaml", &config)
	if err != nil {

		if os.IsNotExist(err) {
			log.Println("config file not found")
			return nil, errors.New("config file not found")
		}

		log.Println("error reading config: %w", err)
		return nil, err
	}

	return &config, nil
}
