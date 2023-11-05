package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	DB struct {
		Port     string `yaml:"port"`
		Host     string `yaml:"host"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	} `yaml:"db"`
}

func LoadConfig() *Config {
	configFile, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		log.Printf("yamlFile get err: %v\n", err)
		os.Exit(1)
	}

	cfg := Config{}
	err = yaml.Unmarshal(configFile, &cfg)
	if err != nil {
		log.Printf("Unmarshal yamlFile err: %v\n", err)
		os.Exit(1)
	}

	return &cfg
}
