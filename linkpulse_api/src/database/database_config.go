package database

import (
	"encoding/json"
	"fmt"
	"os"
)

type DatabaseConfig struct {
	Host         string `json:"host"`
	Port         string `json:"port"`
	User         string `json:"user"`
	Password     string `json:"password"`
	DbName       string `json:"dbname"`
	PoolMaxConns int    `json:"poolMaxConns"`
}

// Loads the database configuration from a JSON file
func LoadConfig(configPath string) (*DatabaseConfig, error) {
	fmt.Println("Loading database configuration file...")
	file, err := os.Open(configPath)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := &DatabaseConfig{}

	if err = decoder.Decode(config); err != nil {
		return nil, err
	}

	fmt.Println("Configuration file loaded successfully!")

	return config, nil
}
