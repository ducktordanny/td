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

func ReadFullConfig() TdConfig {
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

func GetGlobalTodos() []TodoModel {
	config := ReadFullConfig()
	return config.Globals
}

func GetLocalTodos() []TodoModel {
	config := ReadFullConfig()
	return GetLocalModelByPath(&config, GetLocalPath()).Items
}

func WriteFullConfig(conf TdConfig) {
	fileContent, err := json.Marshal(conf)
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
