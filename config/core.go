package config

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

type Configuration struct {
	Port            string `json:"port"`
	DefaultLanguage string `json:"default_language"`
	LegacyEndpoint  string `json:"legacy_endpoint"`
	DatabaseType    string `json:"database_type"`
	DatabaseURL     string `json:"database_url"`
}

func (c *Configuration) LoadFromEnv() {
	if lang := os.Getenv("DEFAULT_LANGUAGE"); lang != "" {
		c.DefaultLanguage = lang
	}
	if port := os.Getenv("PORT"); port != "" {
		c.Port = port
	}
}

func (c *Configuration) LoadFromJSON(path string) error {
	log.Printf("Loading configuration from %s", path)
	b, err := os.ReadFile(path)
	if err != nil {
		log.Printf("Unable to read configuration file %s,%s", path, err)
		return errors.New("unable to read configuration file")

	}
	if err := json.Unmarshal(b, c); err != nil {
		log.Printf("Unable to parse configuration file %s,%s", path, err)
		return errors.New("unable to parse configuration file")
	}
	if c.Port == "" {
		log.Printf("Empty port, reverting to default")
		c.Port = defaultConfiguration.Port
	}
	if c.DefaultLanguage == "" {
		log.Printf("Empty default language, reverting to default")
		c.DefaultLanguage = defaultConfiguration.DefaultLanguage
	}
	return nil
}

var defaultConfiguration = Configuration{
	Port:            ":8080",
	DefaultLanguage: "english",
}
