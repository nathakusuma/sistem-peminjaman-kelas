package config

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/iamolegga/enviper"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Env struct {
	Env               string        `mapstructure:"ENV"`
	AppPort           string        `mapstructure:"APP_PORT"`
	AppURL            string        `mapstructure:"APP_URL"`
	AppName           string        `mapstructure:"APP_NAME"`
	FrontendURL       string        `mapstructure:"FRONTEND_URL"`
	DBHost            string        `mapstructure:"DB_HOST"`
	DBPort            string        `mapstructure:"DB_PORT"`
	DBUser            string        `mapstructure:"DB_USER"`
	DBPass            string        `mapstructure:"DB_PASS"`
	DBName            string        `mapstructure:"DB_NAME"`
	DBSSLMode         string        `mapstructure:"DB_SSLMODE"`
	JwtSecretKey      []byte        // JWT_SECRET_KEY
	JwtExpireDuration time.Duration // JWT_ACCESS_EXPIRE_DURATION
	SMTPHost          string        `mapstructure:"SMTP_HOST"`
	SMTPPort          int           `mapstructure:"SMTP_PORT"`
	SMTPUsername      string        `mapstructure:"SMTP_USERNAME"`
	SMTPEmail         string        `mapstructure:"SMTP_EMAIL"`
	SMTPPassword      string        `mapstructure:"SMTP_PASSWORD"`
}

var (
	viperInstance *viper.Viper
	env           *Env
	once          sync.Once
)

// GetEnv initializes and returns the environment configuration
func GetEnv() *Env {
	once.Do(func() {
		viperInstance = viper.New()
		env = &Env{}

		// Enable environment variables first
		viperInstance.AutomaticEnv()

		// Check if APP_ENV is set in environment variables
		if appEnv := os.Getenv("APP_ENV"); appEnv != "" {
			handleEnvVariables(appEnv, viperInstance, env)
		} else {
			handleEnvFile(viperInstance, env)
		}

		// Process JWT configurations
		handleManuallyParsedVariables(viperInstance, env)

		// Parse durations
		if err := parseDurations(env); err != nil {
			log.Fatal().Msgf("[ENV] failed to parse durations: %s", err.Error())
		}
	})

	return env
}

// handleEnvVariables handles the logic when APP_ENV is set in environment variables
func handleEnvVariables(appEnv string, viperInstance *viper.Viper, env *Env) {
	log.Info().Msgf("[ENV] Using %s environment variables", appEnv)

	// Unmarshal configuration with enviper due to issue with viper
	if err := enviper.New(viperInstance).Unmarshal(env); err != nil {
		log.Fatal().Msgf("[ENV] failed to unmarshal configuration: %s", err.Error())
	}
}

// handleEnvFile handles the logic when APP_ENV is not set in environment variables
func handleEnvFile(viperInstance *viper.Viper, env *Env) {
	if _, err := os.Stat(".env"); err != nil {
		log.Fatal().Msg("[ENV] ENV is not set in environment variables")
		return
	}

	viperInstance.SetConfigFile(".env")
	if err := viperInstance.ReadInConfig(); err != nil {
		log.Fatal().Msg("[ENV] Failed to read .env file")
		return
	}

	log.Info().Msg("[ENV] Using .env file")

	// Unmarshal configuration
	if err := viperInstance.Unmarshal(env); err != nil {
		log.Fatal().Msgf("[ENV] failed to unmarshal configuration: %s", err.Error())
	}
}

// handleManuallyParsedVariables processes the JWT configurations
func handleManuallyParsedVariables(viperInstance *viper.Viper, env *Env) {
	env.JwtSecretKey = []byte(viperInstance.GetString("JWT_ACCESS_SECRET_KEY"))
}

// SetEnv is used in testing to set the environment
func SetEnv(mockEnv *Env) {
	env = mockEnv
}

// Helper function to parse JWT durations
func parseDurations(env *Env) error {
	var err error

	env.JwtExpireDuration, err = time.ParseDuration(viperInstance.GetString("JWT_ACCESS_EXPIRE_DURATION"))
	if err != nil {
		return fmt.Errorf("invalid JWT_ACCESS_EXPIRE_DURATION: %w", err)
	}

	return nil
}
