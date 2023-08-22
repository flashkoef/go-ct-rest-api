package config

type HttpServerConfig struct {
	Port string
}

type Config struct {
	HttpServer HttpServerConfig
}

func New() *Config {
	return &Config{
		HttpServer: HttpServerConfig{
			Port: getEnv("", "8080"),
		},
	}
}

func getEnv(key, defaultVal string) string {
	return defaultVal
}
