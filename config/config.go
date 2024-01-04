package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func New() *Config {
	return &Config{
		HttpServer: HttpServerConfig{
			Port: getEnvValueStringByKey("CI_HTTP_SERVER_PORT", "8080"),
		},
		Ctp: CtpConfig{
			ClientID:     getEnvValueStringByKey("CTP_CLIENT_ID", ""),
			ClientSecret: getEnvValueStringByKey("CTP_CLIENT_SECRET", ""),
			Scopes:       getEnvValueStringByKey("CTP_SCOPES", ""),
			ProjectKey:   getEnvValueStringByKey("CTP_PROJECT_KEY", ""),
		},
	}
}

func getEnvValueStringByKey(key, defaultVal string) string {
	if value, exists := getEnvParamByKey(key); exists {
		return value
	}

	return defaultVal
}

func getEnvParamByKey(key string) (string, bool) {
	readEnvFile()

	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("invalid type assertion while get a environment parameter")
	}

	return value, ok
}

func readEnvFile() {
	envFilePath := getEnvFilePath()
	viper.SetConfigFile(envFilePath)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("error while reading environment file %s", err)
	}
}

func getEnvFilePath() string {
	environment := os.Getenv("ENVIRONMENT")
	if environment == "test" {
		return "../.env"
	}
	return ".env"
}
