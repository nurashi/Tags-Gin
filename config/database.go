package config

import (
	"GinProject1/error_handling"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "7878"
	dbName   = "ginBase"
)

func DatabaseConnection() *gorm.DB {

	sqlInfo := fmt.Sprintf("host = %s port = %d user = %s password = %s dbName = %s sslmode=disable", host, port, user, password, dbName)
	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	error_handling.ErrorPanice(err)

	return db
}
