package config

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Storage  Storage
	LogPath  string `yaml:"logPath" env-default:"../logs/logs.log"`
	LogLevel string `env-default:"info" env:"LOG_LEVEL"`
	Server   Server `yaml:"server"`
}

type Storage struct {
	Path     string
	User     string `env-required:"true" env:"POSTGRES_USER"`
	Password string `env-required:"true" env:"POSTGRES_PASSWORD"`
	DB       string `env-required:"true" env:"POSTGRES_DB"`
	Host     string `env-required:"true" env:"POSTGRES_HOST"`
	Port     string `env-required:"true" env-default:"5432" env:"POSTGRES_PORT"`
}

type Server struct {
	Address     string `yaml:"address" env-default:"80"`
	Timeout     int    `yaml:"timeout" env-default:"4"`
	IdleTimeout int    `yaml:"idleTimeout" env-default:"45"`
	User        string `env-required:"true" env:"URLSHORTENER_USER"`
	Password    string `env-required:"true" env:"URLSHORTENER_PASSWORD"`
}

func MustLoad() (*Config, error) {

	var config Config

	err := cleanenv.ReadConfig("./config.yaml", &config)
	if err != nil {

		if os.IsNotExist(err) {
			log.Println("config file not found")
			return nil, errors.New("config file not found")
		}

		log.Println("error reading config: %w", err)
		return nil, err
	}

	config.buildConnString()

	return &config, nil
}

func (c *Config) buildConnString() {
	c.Storage.Path = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.Storage.User, url.QueryEscape(c.Storage.Password), c.Storage.Host, c.Storage.Port, c.Storage.DB)
}
