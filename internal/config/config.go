package config

import "time"

const EnvPath = "local.env"

type Config struct {
	LogLevel     string        `envconfig:"LOG_LEVEL" default:"info"`
	ListenAddr   string        `envconfig:"LISTEN_ADDR" default:":8080"`
	WriteTimeout time.Duration `envconfig:"WRITE_TIMEOUT" default:"15s"`
	ServerName   string        `envconfig:"SERVER_NAME" default:"ApiServer"`
	Token        string        `envconfig:"TOKEN" required:"true"`
}
