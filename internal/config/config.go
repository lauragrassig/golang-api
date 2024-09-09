package config

import (
	"context"
	"log"

	_ "github.com/joho/godotenv/autoload"
	"github.com/sethvargo/go-envconfig"
)

type Envs struct{
	PostgresHostname string  `env:"POSTGRES_HOSTNAME, required"`
	PostgresPort string `env:"POSTGRES_PORT, required"`
	PostgresUser string `env:"POSTGRES_USER, required"`
	PostgresPassword string `env:"POSTGRES_PASSWORD, required"`
	PostgresDatabase string `env:"POSTGRES_DATABASE, required"`
	ApiPort string `env:"API_PORT, required"`
}

func Init() *Envs {
	ctx := context.Background()

	var e Envs

	err := envconfig.Process(ctx, &e)

	if err != nil {
		log.Fatalf("error loading envs: %v", err)
	}

	return &e
}