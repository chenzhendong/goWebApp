package common

import (
	"encoding/json"
	"os"
	"log"
)

type Config struct {
	Module string `json:"module"`
	Port   int `json:"port"`
}

func Run() {
	log.Println("Reading configuration ...")
	config := readConfig()

	switch config.Module {
	case "auth":
	default:
	}

}

func readConfig() Config {
	var config Config
	file, err := os.Open("appConfig.json")
	if err != nil {
		log.Fatal("Failed to open configuration file", err)
	}
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		log.Fatal("Failed to load configuration. ", err)
	}
	return config
}
