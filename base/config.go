package base

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

const configPath = "./config.json"

var (
	_once   sync.Once
	_config *Config
)

type Config struct {
	Key string `json:"key"`
	URL string `json:"url"`
}

func (c *Config) load() {
	file, err := os.Open(configPath)
	if err != nil {
		log.Fatalf("error open config file: %v", err)
	}
	if err := json.NewDecoder(file).Decode(&c); err != nil {
		log.Fatalf("error parse file: %v", err)
	}

	if c.Key == "" {
		log.Fatal("key required in config file")
	}
	if c.URL == "" {
		log.Fatalf("url required in config file")
	}
}

func GetConfig() *Config {
	_once.Do(func() {
		_config = new(Config)
		_config.load()
	})
	return _config
}
