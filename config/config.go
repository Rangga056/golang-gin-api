package config

import (
	"database/sql"
	"fmt"

	"github.com/Rangga056/golang-gin-api/helper"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbName   = "test"
)

// Check if the database already exists
func databaseExists(sqlInfo, dbName string) (bool, error) {
	db, err := sql.Open("postgres", sqlInfo)
	if err != nil {
		return false, err
	}
	defer db.Close()

	var exists bool
	query := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM pg_database WHERE datname = '%s')", dbName)
	err = db.QueryRow(query).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func DatabaseConnection() *gorm.DB {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, password)

	// Check if the database already exists
	exists, err := databaseExists(sqlInfo, dbName)
	helper.ErrorPanic(err)

	if !exists {
		//Connect to database
		db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
		helper.ErrorPanic(err)

		//create database
		createDBSQL := fmt.Sprintf("CREATE DATABASE %s;", dbName)
		if err := db.Exec(createDBSQL).Error; err != nil {
			helper.ErrorPanic(err)
		}

	}

	// Reconnect to the newly created database
	sqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	helper.ErrorPanic(err)

	return db
}
