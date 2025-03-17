package config

import "time"

const EnvPath = "local.env"

type Config struct {
	LogLevel   string `envconfig:"LOG_LEVEL" default:"info"`
	Rest       Rest
	PostgreSQL PostgreSQL
}

type Rest struct {
	ListenAddr   string        `envconfig:"LISTEN_ADDR" default:":8080"`
	WriteTimeout time.Duration `envconfig:"WRITE_TIMEOUT" default:"15s"`
	ServerName   string        `envconfig:"SERVER_NAME" default:"ApiServer"`
	Token        string        `envconfig:"TOKEN" required:"true"`
}

type PostgreSQL struct {
	Host                string        `envconfig:"DB_HOST" required:"true"`
	Port                int           `envconfig:"DB_PORT" required:"true"`
	Name                string        `envconfig:"DB_NAME" required:"true"`
	User                string        `envconfig:"DB_USER" required:"true"`
	Password            string        `envconfig:"DB_PASSWORD" required:"true"`
	SSLMode             string        `envconfig:"DB_SSL_MODE" default:"disable"`
	PoolSize            int           `envconfig:"DB_POOL_MAX_CONNS" default:"10"`
	PoolConnLifeTime    time.Duration `envconfig:"DB_POOL_MAX_CONN_LIFETIME" default:"180s"`
	PoolMaxConnIdleTime time.Duration `envconfig:"DB_POOL_MAX_CONN_IDLE_TIME" default:"100s"`
}
