package configs

import (
	"flag"
	"os"
)

//type Config struct {
//	Endpoint string
//	DataBase string
//	Accrual  string
//}
//
//type Option func(*Config)
//
//func WithEndpoint(endpoint string) Option {
//	return func(c *Config) {
//		c.Endpoint = endpoint
//	}
//}
//
//func WithDataBase(db string) Option {
//	return func(c *Config) {
//		c.DataBase = db
//	}
//}
//
//func WithAccrual(accrual string) Option {
//	return func(c *Config) {
//		c.Accrual = accrual
//	}
//}
//
//func NewConfig(option ...Option) *Config {
//	cfg := &Config{
//		Endpoint: "8000",
//		DataBase: "postgres://postgres:qwerty@localhost:5434/postgres?sslmode=disable",
//		Accrual:  "http://localhost:8080/api/orders/",
//	}
//
//	for _, opt := range option {
//		opt(cfg)
//	}
//
//	return cfg
//}
//
//func NewServer() *Config {
//	var (
//		endpoint string
//		database string
//		accrual  string
//	)
//	flag.StringVar(&endpoint, "a", "8000", "endpoint to run server")
//	flag.StringVar(&database, "d", "", "db address")
//	flag.StringVar(&accrual, "r", "http://localhost:8080/api/orders/", "accrual")
//
//	flag.Parse()
//
//	if envEndpoint := os.Getenv("RUN_ADDRESS"); endpoint != "" {
//		endpoint = envEndpoint
//	}
//
//	if db := os.Getenv("DATABASE_URI"); db != "" {
//		database = db
//	}
//
//	if envAccrual := os.Getenv("ACCRUAL_SYSTEM_ADDRESS"); accrual != "" {
//		accrual = envAccrual
//	}
//
//	cfg := NewConfig(WithEndpoint(endpoint), WithDataBase(database), WithAccrual(accrual))
//
//	return cfg
//}

type Config struct {
	RunAddr           string
	AccrualSystemAddr string
	DatabaseDSN       string
}

//func newConfig(option options) *Config {
//	cfg := &Config{
//		Endpoint: option.url,
//		DataBase: option.dataBaseDSN,
//		Accrual:  option.accrual,
//	}
//
//	return cfg
//}
//
//type options struct {
//	url         string
//	dataBaseDSN string
//	accrual     string
//}

func ParseServerConfig() *Config {
	cfg := &Config{}

	if cfg.RunAddr = os.Getenv("RUN_ADDRESS"); cfg.RunAddr == "" {
		flag.StringVar(&cfg.RunAddr, "a", "localhost:8080", "Server address")
	}

	if cfg.AccrualSystemAddr = os.Getenv("ACCRUAL_SYSTEM_ADDRESS"); cfg.AccrualSystemAddr == "" {
		flag.StringVar(&cfg.AccrualSystemAddr, "r", "", "Accural system address")
	}

	if cfg.DatabaseDSN = os.Getenv("DATABASE_URI"); cfg.DatabaseDSN == "" {
		flag.StringVar(&cfg.DatabaseDSN, "d", "postgres://postgres:qwerty@localhost:5434/postgres?sslmode=disable", "")
	}

	flag.Parse()

	return cfg
}
