package main

import (
	"log"

	"github.com/devfurkankizmaz/trueties-be/bootstrap"
)

func main() {
	env, err := bootstrap.LoadEnv()
	if err != nil {
		log.Fatalf("Failed to load environment: %v", err)
	}

	// Veritabanı bağlantısını oluştur
	db, err := bootstrap.NewDBConnection(env)
	if err != nil {
		log.Fatalf("Failed to connect to the Database: %v", err)
	}
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("Failed to get SQLDB: %v", err)
		}
		sqlDB.Close()
	}()
}
