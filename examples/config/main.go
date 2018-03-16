package main

import (
	"fmt"
	"log"

	"github.com/andreipimenov/web/config"
)

//Example of using web/config package for application configuration

func main() {
	//Default configuration
	cfg := struct {
		Port int `json:"port"`
	}{
		Port: 9000,
	}

	//Create configuration driver
	driver := config.NewFile("config.json")

	//Get config data and unmarshal as JSON into defined before structure
	err := driver.GetJSON(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Configuration:\n%v\n", cfg)
}
