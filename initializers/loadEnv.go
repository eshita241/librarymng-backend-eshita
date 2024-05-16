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

/*
PostgreSQL is a powerful open-source relational database management system (RDBMS). It is used to store and manage structured data in tables with relationships between them.
The official PostgreSQL Docker image provides a reliable and easy way to deploy PostgreSQL databases as Docker containers. It includes the PostgreSQL server and necessary tools to manage databases.
pgAdmin is a popular open-source web-based administration tool for PostgreSQL. It provides a graphical interface for managing databases, running queries, creating and modifying database objects, and monitoring server activity.
The dpage/pgadmin4 Docker image provides a pre-configured pgAdmin instance that can be easily deployed as a Docker container. It includes pgAdmin along with a web server and other necessary components.

Using Docker images for PostgreSQL and pgAdmin; instead of official offers several benefits:
Consistency:
Docker images provide consistent environments across different systems. You can ensure that the same versions of PostgreSQL and pgAdmin are used in development, testing, and production environments.
Isolation:
Docker containers provide isolation for applications and services. Each container runs in its own environment, preventing conflicts with other software installed on the host system.
Portability:
Docker images can be easily shared and distributed. You can package your application along with its dependencies (including PostgreSQL and pgAdmin) into a single Docker image, making it portable across different platforms and environments.
Ease of Deployment:
Docker simplifies the deployment process by providing a consistent way to package, distribute, and run applications. With Docker Compose, you can define multi-container applications and manage them with a single command.
Resource Efficiency:
Docker containers are lightweight and consume fewer resources compared to virtual machines. They can be quickly started, stopped, and scaled as needed, making efficient use of system resources.
*/
