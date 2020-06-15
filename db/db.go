package db

import (
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

// DBClient is a instance of the database client
type DBClient struct {
	db *gorm.DB
}

// NewDBClient returns a new instance of a database client
func NewTargetDBClient() (DBClient, error) {
	var result DBClient

	// Load ENVs
	err := godotenv.Load("../.env")
	if err != nil {
		err = godotenv.Load(".env")
		if err != nil {
			return result, err
		}
	}

	mysqlDB := os.Getenv("MYSQL_DATABASE")
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlPort := os.Getenv("MYSQL_PORT")

	// Create connection string
	connectionString := mysqlUser + ":" + mysqlPassword + "@tcp(" + mysqlHost + ":" + mysqlPort + ")/" + mysqlDB + "?charset=utf8&parseTime=True&loc=Local"

	// Open connection
	result.db, err = gorm.Open("mysql", connectionString)
	if err != nil {
		return result, err
	}

	result.db.SingularTable(true)

	// Migrate the schema
	result.db.AutoMigrate(&Products{})
	result.db.AutoMigrate(&Media{})
	result.db.AutoMigrate(&ProductAtt{})

	result.db.DB().SetMaxIdleConns(0)
	result.db.DB().SetConnMaxLifetime(1 * time.Second)

	// Enable Logger, show detailed log
	// result.db.LogMode(true)

	return result, err
}

// CloseConnection closes the client's connection
// It returns an error if one occurred
func (dbc *DBClient) CloseConnection() error {
	return dbc.db.Close()
}
