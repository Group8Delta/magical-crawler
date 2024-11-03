package config

import (
	"log"
	"time"

	"sync"

	"github.com/spf13/viper"
)

var once sync.Once
var config *Config

type Config struct {
	DatabaseHost            string        `mapstructure:"DATABASE_HOST"`
	DatabasePort            string        `mapstructure:"DATABASE_PORT"`
	DatabaseUser            string        `mapstructure:"DATABASE_USER"`
	DatabasePassword        string        `mapstructure:"DATABASE_PASSWORD"`
	DatabaseName            string        `mapstructure:"DATABASE_NAME"`
	DatabaseSSLMode         string        `mapstructure:"DATABASE_SSLMODE"`
	DatabaseMaxIdleConns    int           `mapstructure:"DATABASE_MAX_IDLE_CONNS"`
	DatabaseMaxOpenConns    int           `mapstructure:"DATABASE_MAX_OPEN_CONNS"`
	DatabaseConnMaxLifetime time.Duration `mapstructure:"DATABASE_CONN_MAX_LIFETIME"`
}

func loadEnvConfig() (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigFile("../../.env")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		log.Printf("Unable to read .env file: %v", err)
		return nil, err
	}
	return v, nil
}

func parseConfig(v *viper.Viper) (*Config, error) {
	var cfg Config
	err := v.Unmarshal(&cfg)
	if err != nil {
		log.Printf("Unable to parse config: %v", err)
		return nil, err
	}

	return &cfg, nil
}

func GetConfig() *Config {

	once.Do(func() {
		v, err := loadEnvConfig()
		if err != nil {
			log.Fatalf("Error loading environment config: %v", err)
		}

		config, err = parseConfig(v)
		if err != nil {
			log.Fatalf("Error parsing config: %v", err)
		}

	})

	return config
}
