package main

import (
	"log"

	"github.com/JoaoGSantiago/starti-backend/internal/config"
	"github.com/JoaoGSantiago/starti-backend/internal/db"
	"github.com/JoaoGSantiago/starti-backend/internal/handlers"
	"github.com/JoaoGSantiago/starti-backend/internal/repositories"
	"github.com/JoaoGSantiago/starti-backend/internal/router"
	"github.com/JoaoGSantiago/starti-backend/internal/services"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("sem .env")
	}

	cfg := config.Load()
	db := db.Connect(cfg)

	// camada de repositorio
	userRepo := repositories.NewUserRepository(db)
	postRepo := repositories.NewPostRepository(db)
	commentRepo := repositories.NewCommentRepository(db)

	// camada de serviços
	jwtSvc := services.NewJWTService(cfg.JWTSecret)
	authSvc := services.NewAuthService(userRepo, jwtSvc)
	userSvc := services.NewUserService(userRepo)
	postSvc := services.NewPostService(postRepo)
	commentSvc := services.NewCommentService(commentRepo)

	// camada de handlers
	authHandler := handlers.NewAuthHandler(authSvc)
	userHandler := handlers.NewUserHandler(userSvc)
	postHandler := handlers.NewPostHandler(postSvc)
	commentHandler := handlers.NewCommentHandler(commentSvc)

	r := router.Setup(jwtSvc, authHandler, userHandler, postHandler, commentHandler)

	log.Printf("server deu run em :%s", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("server falhou ao iniciar: %v", err)
	}
}
