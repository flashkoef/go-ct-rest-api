package config

type HttpServerConfig struct {
	Port string
}

type CtpConfig struct {
	ClientID     string
	ClientSecret string
	Scopes       string
	ProjectKey   string
}

type Config struct {
	HttpServer HttpServerConfig
	Ctp        CtpConfig
}
