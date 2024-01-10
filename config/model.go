package config

type HTTPServerConfig struct {
	Port string
}

type CtpConfig struct {
	ClientID     string
	ClientSecret string
	Scopes       string
	ProjectKey   string
}

type Config struct {
	HTTPServer HTTPServerConfig
	Ctp        CtpConfig
}
