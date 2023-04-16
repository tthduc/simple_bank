package util

import (
	"bytes"
	"github.com/spf13/viper"
	"os"
)

// Config stores all configuration of the application
// The values are read by viper from a config file or environment variables
type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER" yaml:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE" yaml:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS" yaml:"SERVER_ADDRESS"`
}

var (
	Instance *Config
)

// loadConfig reads configuration from file or environment variables
func LoadConfig(path string) (config Config, err error) {
	buf, err := os.ReadFile(path)
	if err != nil {
		return
	}

	/* assign config value to cfg */
	viper.SetConfigType("yaml")
	err = viper.ReadConfig(bytes.NewBuffer(buf))
	if err != nil {
		return
	}

	/* overwrite with the value of EVN variable by matching the uppercased key */
	viper.AutomaticEnv()

	/* binding to cfg */
	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}

	Instance = &config

	return
}
