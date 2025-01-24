package postgres

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBOptions struct {
	Host     string
	port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func (options *DBOptions) PostgresDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", options.Host, options.port, options.User, options.Password, options.DBName, options.SSLMode)
}

func NewDB(options *DBOptions) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(options.PostgresDSN()), &gorm.Config{})
	// if err !=nil{
	// 	panic(fmt.Sprintf("Failed to connect to database: %v", err))
	// }
}
