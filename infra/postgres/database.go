package infra

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDB() (*gorm.DB, error) {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USERNAME")
	pass := os.Getenv("DB_PASSWORD")

	dbConf := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, name, pass)
	fmt.Println(dbConf)

	db, err := gorm.Open(postgres.Open(dbConf), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
