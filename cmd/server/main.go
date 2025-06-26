package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/ldcmleo/Ichtaca/internal/config"
	"github.com/ldcmleo/Ichtaca/internal/storage"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("[ERROR] cannot load .env file: ", err.Error())
	}

	cfg := config.Load()

	db, err := storage.NewDB(cfg)
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}

	if err := storage.Migrate(cfg, db); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

}
