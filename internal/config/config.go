package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Port       string `yaml:"port" env:"PORT" env-default:"8080"`
	Host       string `yaml:"host" env:"HOST" env-default:"localhost"`
	DBPassword string `env:"TAG_DB_PASSWORD" env-default:"test_pass"`
	DBHost     string `env:"TAG_DB_HOST" env-default:"127.0.0.1:3036"`
}

func MustLoad() *Config {

	var cfg Config

	// Читаем конфиг-файл и заполняем нашу структуру
	err := cleanenv.ReadConfig("./.env", &cfg)
	if err != nil {
		log.Fatalf("error reading config file: %s", err)
	}

	return &cfg
}
