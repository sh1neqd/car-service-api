package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type Config struct {
	PostgresSQL struct {
		Host string `env:"PSQL_HOST"  env-default:"db"`
		//Host     string `env:"PSQL_HOST"  env-default:"localhost"`
		Port     int    `env:"PSQL_PORT"  env-default:"5432"`
		Username string `env:"PSQL_USERNAME"  env-default:"postgres"`
		Password string `env:"PSQL_PASSWORD" env-default:"postgres"`
		Database string `env:"PSQL_DATABASE"  env-default:"postgres"`
		//Database string `env:"PSQL_DATABASE"  env-default:"carsservice"`
		SSLMode string `env:"SSL_MODE" env-default:"disable"`
	}
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		log.Print("gather config")

		instance = &Config{}

		if err := cleanenv.ReadEnv(instance); err != nil {
			helpText := "sh1neqd - car service"
			help, _ := cleanenv.GetDescription(instance, &helpText)
			log.Print(help)
			log.Fatal(err)
		}
	})
	return instance
}
