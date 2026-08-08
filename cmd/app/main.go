package main

import (
	srv "Geoapi/cmd/server"
	log "Geoapi/internal/logger"
)

var logger = log.Logger()

func main() {
	srv, err := srv.New()
	if err != nil{
		logger.Error().Err(err).Msg("Failed to start server")
	}
	if err := srv.Start(); err != nil{
		logger.Error().Err(err).Msg("Failed to start server")
	}else{
		logger.Info().Msg("Starting App")
	}
}
