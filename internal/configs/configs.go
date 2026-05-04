package configs

import (
	"Geoapi/internal/logger"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var logs = logger.Logger()

type Configs struct{
	ServerPort 	string `envconfig:"SERVER_PORT" required:"true"`
	APIkey 	string `envconfig:"API_KEY" required:"true"`
}

func Load() (*Configs, error){
	var cfg Configs
	if err := godotenv.Load(); err != nil{
		logs.Error().Err(err).Msg("Failed to get .env file")
		return nil, err
	}
	if err := envconfig.Process("", &cfg); err != nil{
		logs.Error().Err(err).Msg("Failed to get configs")
		return nil, err
	}
	return &cfg, nil
}