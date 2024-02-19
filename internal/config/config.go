package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"time"
)

type Config struct {
	HTTPServer `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func MustLoad() Config {

	configPath := "config/local.yaml"

	var cfg Config

	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		log.Printf("can't open config file, error: %v", err)
	}
	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		log.Fatalf("error unmarshal yaml file: %v", err)
	}

	return cfg
}
