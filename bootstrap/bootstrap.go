package bootstrap

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Env struct {
	DBHost string `mapstructure:"POSTGRES_HOST"`
	DBUser string `mapstructure:"POSTGRES_USER"`
	DBPass string `mapstructure:"POSTGRES_PASSWORD"`
	DBName string `mapstructure:"POSTGRES_DB"`
	DBPort string `mapstructure:"POSTGRES_PORT"`
}

func LoadEnv() (*Env, error) {
	env := &Env{}
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(env); err != nil {
		return nil, err
	}

	return env, nil
}

func NewDBConnection(config *Env) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.DBHost, config.DBUser, config.DBPass, config.DBName, config.DBPort)

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	DB.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Connected to DB")
	return DB, nil
}
