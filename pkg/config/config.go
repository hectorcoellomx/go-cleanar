package config

import "github.com/spf13/viper"

type Config struct {
	DatabaseURL string
	// Otras configuraciones
}

func LoadConfig() *Config {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	return &Config{
		DatabaseURL: viper.GetString("DATABASE_URL"),
		// Configuraciones adicionales
	}
}

func InitDatabase(databaseURL string) *gorm.DB {
	// Lógica para inicializar y configurar la conexión a la base de datos utilizando GORM
}
