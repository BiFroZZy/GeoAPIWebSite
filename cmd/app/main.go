package main

import (
	srv "Geoapi/cmd/server"
	"Geoapi/internal/logger"
)

func main() {
	srv, err := srv.New()
	var logger = logger.Logger()

	if err != nil{
		logger.Error().Err(err).Msg("Failed to start server")
	}
	if err := srv.Start(); err != nil{
		logger.Error().Err(err).Msg("Failed to start server")
	}
}
