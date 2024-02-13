package config

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"
)

func GetConfigPath() string {
	const configBase = ".config/.tdconfig.json"

	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln(err)
	}

	return filepath.Join(home, configBase)
}

// TODO: If there is no config file then we should create a default one, or have it in the install process?
func ReadConfig() TdConfig {
	jsonConfig, err := os.Open(GetConfigPath())

	if err != nil {
		log.Fatalln(err)
	}
	defer jsonConfig.Close()

	var config TdConfig
	byteValue, err := io.ReadAll(jsonConfig)
	json.Unmarshal([]byte(byteValue), &config)

	return config
}
