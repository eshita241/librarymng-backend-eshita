package database

import (
	"fmt"
	"librarymng-backend/initializers"
	"librarymng-backend/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB // DbInstance is a representation of the database instance - its only field is the gorm DB instance
}

func GlobalActivationScope(db *gorm.DB) *gorm.DB {
	return db.Where("is_activated = ?", true) //The GlobalActivationScope function in Go, using the GORM ORM library, defines a query scope that filters records in the database to only include those where the is_activated field is true.
}

var Database DbInstance

func ConnectToDB(config *initializers.Config) { //takes Config struct
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}) //dsn in form
	if err != nil {
		log.Fatal("error connecting to database: %w", err)
	}

	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	//This line executes a raw SQL statement to create the uuid-ossp extension in the PostgreSQL database if it does not already exist. The uuid-ossp extension provides functions to generate universally unique identifiers (UUIDs).
	db.Logger = logger.Default.LogMode(logger.Info)
	//This line configures GORM to use its default logger in Info mode. This means that GORM will log detailed information about the database operations it performs, such as executed SQL queries, which can be useful for debugging and monitoring.

	if os.Getenv("SHOULD_MIGRATE") == "TRUE" { //perform database migrations based on an environment variable,
		log.Println("Running DB Migrations...")
		if err := db.AutoMigrate(&models.Book{}, &models.User{}, &models.Issue{}); err != nil { //The AutoMigrate method is called with the models Book, User, and Issue. This method automatically creates or updates the database tables to match the provided model definitions. It ensures that the schema is up-to-date with the current state of the application models.
			log.Println("Error running DB Migrations:", err)
		} else {
			log.Println("DB Migrations completed")
		}
	}

	Database = DbInstance{Db: db}
}

/*
	A DSN (Data Source Name) is a string that provides the necessary information to connect to a database. It typically includes:
	Username
	Password
	Host (server address)
	Port
	Database name
	Optional parameters (like SSL mode, timeouts, etc.)

	connectionString is a more generic term used in many database connectivity libraries, including ADO.NET, JDBC, and many others.
	It is a string that contains all the necessary information to connect to a database.
	 Unlike a DSN, a connection string is typically specified directly in the application's code or configuration files and includes all the connection details inline.
*/
