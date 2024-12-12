package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Bot     *Bot    `yaml:"bot"`
	Storage Storage `yaml:"storage"`
}

type Bot struct {
	Token       string `yaml:"token"`
	HostAddr    string `yaml:"host_addr"`
	WebhookAddr string `yaml:"webhook_addr"`
}

type Storage struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"db_name"`
}

func MustLoad() *Config {
	file, err := os.Open("config.yaml")
	if err != nil {
		log.Fatalf("failed to open config file: %v", err)
	}
	defer file.Close()

	var config Config
	err = yaml.NewDecoder(file).Decode(&config)
	if err != nil {
		log.Fatalf("failed to decode config file: %v", err)
	}

	return &config
}
