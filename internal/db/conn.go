package db

import (
	"log"

	"github.com/JoaoGSantiago/starti-backend/internal/config"
	"github.com/JoaoGSantiago/starti-backend/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect(cfg *config.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("falha ao conectar no banco de dados: %v", err)
	}

	// AutoMigrate cria ou atualiza as tabelas sem apagar dados existentes
	if err := db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{}); err != nil {
		log.Fatalf("falha ao rodar migrations: %v", err)
	}

	log.Println("banco de dados conectado e migration rodado")
	return db
}
