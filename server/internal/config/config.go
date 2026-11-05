package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type DatabaseConfig struct {
	Host string `yaml:"host"`
	Port uint16 `yaml:"port"`
}

type LoggingConfig struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
}

type ServerConfig struct {
	Port uint16 `yaml:"port"`
}

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	Logger   LoggingConfig  `yaml:"logger"`
}

func LoadConfig() Config {
	var config Config
	err := cleanenv.ReadConfig("config.yaml", &config)
	if err != nil {
		log.Fatalln("Error while loading app configuration")
	}

	return config
}
