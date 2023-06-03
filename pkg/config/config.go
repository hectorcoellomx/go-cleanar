package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DB_Host     string
	DB_Port     string
	DB_User     string
	DB_Password string
	DB_Name     string
}

func LoadConfig() *Config {

	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	return &Config{
		DB_Host:     os.Getenv("DB_SERVER"),
		DB_Port:     os.Getenv("DB_PORT"),
		DB_User:     os.Getenv("DB_USER"),
		DB_Password: os.Getenv("DB_PASSWORD"),
		DB_Name:     os.Getenv("DB_DATABASE"),
	}
}

func InitDatabase(conf *Config) (*gorm.DB, error) {
	dsn := conf.DB_User + ":" + conf.DB_Password + "@tcp(" + conf.DB_Host + ":" + conf.DB_Port + ")/" + conf.DB_Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
