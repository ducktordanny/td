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

	usr, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln(err)
	}

	return filepath.Join(usr, configBase)
}

func GetLocalPath() string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	return path
}

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

func WriteConfig(conf TdConfig) {
	fileContent, err := json.MarshalIndent(conf, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}
	jsonConfig, err := os.Create(GetConfigPath())
	if err != nil {
		log.Fatalln(err)
	}
	defer jsonConfig.Close()
	io.WriteString(jsonConfig, string(fileContent))
}
