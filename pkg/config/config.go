package config

import (
	"log"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Config struct {
	DatabaseURL string
	// Otras configuraciones
}

func LoadConfig() *Config {

	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	return &Config{
		DatabaseURL: viper.GetString("DATABASE_URL"),
		// Configuraciones adicionales
	}
}

func InitDatabase(databaseURL string) *gorm.DB {
	// Lógica para inicializar y configurar la conexión a la base de datos utilizando GORM
}
