package main

import (
	"api/oportunities/config"
	"api/oportunities/router"
)

var (
	logger config.Logger
)

func main () {

	logger = *config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization err: %v", err)
		return;
	}


	router.Initialize()
}