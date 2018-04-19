package main

import (
	"fmt"
	"log"

	"github.com/andreipimenov/goto/config"
)

// Example of using config package for application configuration.
func main() {
	// Default configuration.
	cfg := struct {
		Port int `json:"port"`
	}{
		Port: 9000,
	}

	// Create configuration driver.
	driver := config.NewFileConfig("config.json")

	// Get config data and unmarshal as JSON into structure defined before.
	err := driver.GetJSON(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Print configuation.
	fmt.Printf("Configuration:\n%v\n", cfg)
}
