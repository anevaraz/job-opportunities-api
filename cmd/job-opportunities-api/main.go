package main

import (
	router "github.com/anevaraz/job-opportunities-api/internal/routes"
	config "github.com/anevaraz/job-opportunities-api/pkg"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}
