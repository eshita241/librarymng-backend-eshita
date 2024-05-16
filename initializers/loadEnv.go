package initializers

import (
	"time" // Importing the time package for time.Duration

	"github.com/spf13/viper" // Importing viper package for configuration management
)

// Config struct defines the structure of our configuration settings
type Config struct {
	DBHost         string `mapstructure:"POSTGRES_HOST"`     // Database host
	DBUserName     string `mapstructure:"POSTGRES_USER"`     // Database username
	DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"` // Database password
	DBName         string `mapstructure:"POSTGRES_DB"`       // Database name
	DBPort         string `mapstructure:"POSTGRES_PORT"`     // Database port

	JwtSecret    string        `mapstructure:"JWT_SECRET"`     // JWT secret key
	JwtExpiresIn time.Duration `mapstructure:"JWT_EXPIRED_IN"` // JWT expiration duration
	JwtMaxAge    int           `mapstructure:"JWT_MAXAGE"`     // JWT max age

	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"` // Client origin
}

// LoadConfig function reads the configuration from a .env file and/or environment variables
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)   // Specify the path to look for the config file
	viper.SetConfigType("env")  // Specify that the config file is of type env
	viper.SetConfigName(".env") // Name of the config file to look for

	viper.AutomaticEnv() // Automatically override config values with environment variables if they exist

	err = viper.ReadInConfig() // Read the config file
	if err != nil {
		return // Return the error if reading the config file fails
	}

	err = viper.Unmarshal(&config) // Unmarshal the config values into the Config struct
	return                         // Return the config struct and any error that occurred
}
