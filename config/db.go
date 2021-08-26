package config

import (
	"os"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     "5432",
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		DBName:   os.Getenv("DB_NAME"),
	}
	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	return "host=" + dbConfig.Host + " port=" + dbConfig.Port + " user=" + dbConfig.User + " dbname=" + dbConfig.DBName + " sslmode=disable password=" + dbConfig.Password
}
