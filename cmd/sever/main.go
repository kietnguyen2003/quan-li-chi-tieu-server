package main

import (
	"log"
	"quan-li-chi-tieu/cmd/config"
	authApp "quan-li-chi-tieu/src/application/auth"
	"quan-li-chi-tieu/src/infrastructure/database"
	"quan-li-chi-tieu/src/infrastructure/security"
	authH "quan-li-chi-tieu/src/interface/http/auth"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	log.Println("Connect database successfully")

	if err := database.AutoMigrate(db); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	log.Println("Database migrated successfully")

	userRepo := database.NewGormUserRepository(db)
	passwordHasher := security.NewBcryptHasher()
	tokenProvider := security.NewJWTProvider(cfg.JWTSecret)

	authUseCase := authApp.NewAuthUseCase(userRepo, passwordHasher, tokenProvider)

	authHandler := authH.NewAuthHandler(authUseCase)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:3001"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))
	setupRoutes(r, authHandler)

	log.Printf("Server starting on port %s...", cfg.Port)
	log.Fatal(r.Run(":" + cfg.Port))

}

func setupRoutes(r *gin.Engine, authHandler *authH.AuthHandler) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/login", authHandler.Login)
	r.POST("/register", authHandler.Register)
}
