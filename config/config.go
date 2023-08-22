package config

import (
	"log"

	"github.com/spf13/viper"
)

type HttpServerConfig struct {
	Port string
}

type Config struct {
	HttpServer HttpServerConfig
}

func New() *Config {
	return &Config{
		HttpServer: HttpServerConfig{
			Port: getEnv("CI_HTTP_SERVER_PORT", "8080"),
		},
	}
}

func getEnv(key, defaultVal string) string {
	if value, exists := getConfigParam(key); exists {
		return value
	}

	return defaultVal
}

func getConfigParam(key string) (string, bool) {
	readConfigFile()
	
	value, ok := viper.Get(key).(string)
	if !ok {
    log.Fatalf("Invalid type assertion while get a config parameter")
  }
	
	return value, ok
}

func readConfigFile() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
}
