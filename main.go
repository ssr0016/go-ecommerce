package main

import (
	"go-ecommerce/config"
	"go-ecommerce/internal/api"
	"log"
)

func main() {

	cfg, err := config.SetUpEnv()
	if err != nil {
		log.Fatalf("config file is not loaded properly %v\n", err)
	}

	api.StartServer(cfg)
}
