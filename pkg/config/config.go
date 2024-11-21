package config

import (
	"github.com/notneet/go-hoyo-daily/pkg/env"

	"github.com/joho/godotenv"
)

type Config struct {
	Host       string
	Port       int
	Env        string
	JWTSecret  string
	JWTExpires int

	Dsn       string
	SentryDsn string
}

func LoadConfig() (*Config, error) {
	cfg := &Config{}
	cfg.Env = env.GetString("ENV", "production")

	if cfg.Env == "development" {
		if err := godotenv.Load(); err != nil {
			return nil, err
		}
	}

	cfg.Host = env.GetString("HOST", "localhost")
	cfg.Port = env.GetInt("PORT", 4444)
	cfg.JWTSecret = env.GetString("JWT_SECRET", "secret")
	cfg.JWTExpires = env.GetInt("JWT_EXPIRES", 3600)

	cfg.Dsn = env.GetString("DSN", "root:root@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local")
	cfg.SentryDsn = env.GetString("SENTRY_DSN", "")

	// if cfg.JWTSecret == "" {
	// 	return nil, errors.New("JWT_SECRET is empty")
	// }

	return cfg, nil
}
