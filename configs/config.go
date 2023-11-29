package configs

import (
	"flag"
	"os"
)

type Config struct {
	Endpoint string
	DataBase string
	Accrual  string
}

type Option func(*Config)

func WithEndpoint(endpoint string) Option {
	return func(c *Config) {
		c.Endpoint = endpoint
	}
}

func WithDataBase(db string) Option {
	return func(c *Config) {
		c.DataBase = db
	}
}

func WithAccrual(accrual string) Option {
	return func(c *Config) {
		c.Accrual = accrual
	}
}

func NewConfig(option ...Option) *Config {
	cfg := &Config{
		Endpoint: "8000",
		DataBase: "postgres://postgres:qwerty@localhost:5434/postgres?sslmode=disable",
		Accrual:  "http://localhost:8080/",
	}

	for _, opt := range option {
		opt(cfg)
	}

	return cfg
}

func NewServer() *Config {
	var cfg Config
	flag.StringVar(&cfg.Endpoint, "a", "8000", "endpoint to run server")
	flag.StringVar(&cfg.DataBase, "d", "", "db address")
	flag.StringVar(&cfg.Accrual, "r", "http://localhost:8080/", "accrual")

	flag.Parse()

	if endpoint := os.Getenv("RUN_ADDRESS"); endpoint != "" {
		cfg.Endpoint = endpoint
	}

	if db := os.Getenv("DATABASE_URI"); db != "" {
		cfg.Endpoint = db
	}

	if accrual := os.Getenv("ACCRUAL_SYSTEM_ADDRESS"); accrual != "" {
		cfg.Endpoint = accrual
	}

	return &cfg
}
