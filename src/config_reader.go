package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig
	Redis  RedisConfig
}

type ServerConfig struct {
	Port        uint32
	Baseurl     string
	Logfile     string
	Storagetype string
}

type RedisConfig struct {
	Ip   string
	Port uint32
}

func (configuration *Config) Init() error {
	viper.SetConfigType("yml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("./../config")

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

	if config.Server.Storagetype == "redis" {
		config.Redis.Ip = "localhost"
		config.Redis.Port = 6379
	}
}

func (config *Config) GetStorageType() string {
	return config.Server.Storagetype
}
