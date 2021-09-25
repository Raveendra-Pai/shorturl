package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig
}

type ServerConfig struct {
	Port    uint32
	Baseurl string
	Logfile string
	Storagetype string
}

func (configuration *Config) Init() error {
	viper.SetConfigType("yml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("./../config")
	//viper.AddConfigPath("C:\\Work\\Go\\shorturl\\config")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
		return err
	}

	setDefaults(configuration)

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
		return err
	}

	return nil
}

func setDefaults(config *Config) {
	config.Server.Baseurl = "localhost"
	config.Server.Port = 8081
	config.Server.Logfile = "shorturl.log"
	config.Server.Storagetype = "inmemory"
}

func (config *Config) GetStorageType() string {
	return config.Server.Storagetype
}
