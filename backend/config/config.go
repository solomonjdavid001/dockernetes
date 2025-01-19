package config

import (
	"bytes"
	"embed"

	"github.com/spf13/viper"
)

// Structs for configuration
type AppConfig struct {
	Name string `mapstructure:"name"`
	Env  string `mapstructure:"env"`
	Port string `mapstructure:"port"`
}

type RedisConfig struct {
	Server string `mapstructure:"server"`
}

type Config struct {
	App   AppConfig   `mapstructure:"app"`
	Redis RedisConfig `mapstructure:"redis"`
}

var GlobalConfig Config

func LoadConfig(configFile embed.FS) error {
	configData, err := configFile.ReadFile("config.yaml")
	if err != nil {
		return err
	}

	viper.SetConfigType("yaml")

	err = viper.ReadConfig(bytes.NewReader(configData))
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&GlobalConfig)
	if err != nil {
		return err
	}
	return nil
}
