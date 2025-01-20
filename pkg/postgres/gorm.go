package postgres

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBOptions struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
	sslmode  string
}

func (options *DBOptions) PostgresDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", options.host, options.port, options.user, options.password, options.dbname, options.sslmode)
}

func NewDB(options *DBOptions) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(options.PostgresDSN()), &gorm.Config{})
	// if err !=nil{
	// 	panic(fmt.Sprintf("Failed to connect to database: %v", err))
	// }
}
