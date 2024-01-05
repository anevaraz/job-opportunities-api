package main

import (
	"github.com/anevaraz/job-opportunities-api/config"
	router "github.com/anevaraz/job-opportunities-api/internal/routes"
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
