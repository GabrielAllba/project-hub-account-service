package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "project-hub/account-service/docs"
	"project-hub/account-service/internal/controller"
	userrouter "project-hub/account-service/internal/delivery/http/router"
	"project-hub/account-service/internal/repository"
	"project-hub/account-service/internal/usecase"
	"project-hub/account-service/pkg/config"

	"project-hub/account-service/pkg/utils/dbutil"
)

// @title           Project Hub Authentication API
// @version         1.0
// @description     A Project Hub Clean Architecture authentication API with PostgreSQL and Gin.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.example.com/support
// @contact.email  support@example.com

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8080
// @BasePath  /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	cfg := config.NewConfig()

	if err := dbutil.EnsureDatabaseExists(cfg.DBDriver, cfg.DBName, cfg.DefaultDBSource); err != nil {
		log.Fatalf("Database setup error: %v", err)
	}

	db, err := sql.Open(cfg.DBDriver, cfg.DBSource)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	if err := dbutil.RunMigrationsFromFile(db, "migrations/schema.sql"); err != nil {
		log.Fatalf("Migration error: %v", err)
	}

	// Intialize the blacklisted token repository
	blacklistedTokenRepo := repository.NewBlacklistedTokenRepository(db)
	blacklistedTokenUseCase := usecase.NewBlacklistedTokenUseCase(blacklistedTokenRepo)

	// Initialize repository, usecase, controller, etc.
	userRepo := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userController := controller.NewUserController(userUseCase, blacklistedTokenUseCase, cfg)

	router := gin.Default()

	userrouter.NewUserRouter(router, userController, cfg, userRepo, blacklistedTokenRepo)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Printf("Server started on %s", cfg.ServerAddr)
	log.Fatal(router.Run(cfg.ServerAddr))
}
