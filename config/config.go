package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
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
	BotToken                string        `mapstructure:"BOT_TOKEN"`
	Port                    string        `mapstructure:"PORT"`
	SheypoorToken           string        `mapstructure:"SHEYPOOR_TOKEN"`
	DivarToken              string        `mapstructure:"DIVAR_TOKEN"`
}

func loadEnvConfig() (*viper.Viper, error) {
	_, b, _, _ := runtime.Caller(0)
	root := filepath.Join(filepath.Dir(b), "..")
	v := viper.New()
	if os.Getenv("DOCKER-DEPLOY") == "" {
		v.SetConfigFile(filepath.Join(root, ".env"))
		v.AutomaticEnv()

		if err := v.ReadInConfig(); err != nil {
			return nil, fmt.Errorf("failed to read config file: %w", err)
		}
	} else {
		bindEnvVars(v)
	}

	// Optional: Check if required variables are set and return an error if any are missing
	missingVars := checkRequiredVars(v, []string{
		"DATABASE_HOST", "DATABASE_PORT", "DATABASE_USER", "DATABASE_PASSWORD",
		"DATABASE_NAME", "DATABASE_SSLMODE", "DATABASE_MAX_IDLE_CONNS",
		"DATABASE_MAX_OPEN_CONNS", "DATABASE_CONN_MAX_LIFETIME", "PORT",
	})
	if len(missingVars) > 0 {
		return nil, fmt.Errorf("missing required environment variables: %v", missingVars)
	}

	return v, nil
}

func bindEnvVars(v *viper.Viper) {
	v.BindEnv("DATABASE_HOST")
	v.BindEnv("DATABASE_PORT")
	v.BindEnv("DATABASE_USER")
	v.BindEnv("DATABASE_PASSWORD")
	v.BindEnv("DATABASE_NAME")
	v.BindEnv("DATABASE_SSLMODE")
	v.BindEnv("DATABASE_MAX_IDLE_CONNS")
	v.BindEnv("DATABASE_MAX_OPEN_CONNS")
	v.BindEnv("DATABASE_CONN_MAX_LIFETIME")
	v.BindEnv("PORT")
	v.BindEnv("SHEYPOOR_TOKEN")
	v.BindEnv("DIVAR_TOKEN")

}

func checkRequiredVars(v *viper.Viper, keys []string) []string {
	var missing []string
	for _, key := range keys {
		if !v.IsSet(key) {
			missing = append(missing, key)
		}
	}
	return missing
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
