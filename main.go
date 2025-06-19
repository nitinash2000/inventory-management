package main

import (
	"encoding/json"
	"fmt"
	"inventory-management/config"
	"inventory-management/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func GetConfig(filePath string) (*config.Config, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config config.Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func main() {
	env := os.Getenv("env")
	if env == "" {
		env = "default"
	}

	config, err := GetConfig(fmt.Sprintf("%s.json", env))
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
		os.Exit(1)
	}

	if config.ServerPort == "" {
		config.ServerPort = "8080"
	}

	r := gin.Default()

	routes.Router(r)

	r.Run(":" + config.ServerPort)
}
