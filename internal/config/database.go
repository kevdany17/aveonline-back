package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
	"os"
	"log"
	"fmt"
)


func InitalizeDataBase() *gorm.DB{
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil
	}
	dsn := "host="+os.Getenv("DATABASE_HOST")+" user="+os.Getenv("DATABASE_USER")+" password="+os.Getenv("DATABASE_PASSWORD")+" dbname="+os.Getenv("DATABASE_NAME")+" port="+os.Getenv("DATABASE_PORT")+" sslmode=disable"
	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Fatal(err)
	}
	return db
}