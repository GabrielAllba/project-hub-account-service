package router

import (
	"project-hub/account-service/internal/controller"
	userhttp "project-hub/account-service/internal/delivery/http/handler/user"
	"project-hub/account-service/internal/delivery/middleware"
	"project-hub/account-service/internal/domain/entity"
	"project-hub/account-service/pkg/config"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func NewUserRouter(
	router *gin.Engine,
	controller controller.UserController,
	cfg *config.Config,
	userRepo entity.UserRepository,
	blacklistedTokenRepo entity.BlacklistedTokenRepository,
) {

	handler := &userhttp.UserHandler{
		UserController: controller,
		Validator:      validator.New(),
	}

	authRoutes := router.Group("/api/auth")
	{
		authRoutes.POST("/register", userhttp.RegisterHandler(handler))
		authRoutes.POST("/login", userhttp.LoginHandler(handler))
	}

	authRoutes.Use(middleware.AuthMiddleware(cfg, userRepo, blacklistedTokenRepo))
	{
		authRoutes.POST("/logout", userhttp.LogoutHandler(handler))
	}

	userRoutes := router.Group("/api/users")
	userRoutes.Use(middleware.AuthMiddleware(cfg, userRepo, blacklistedTokenRepo))
	{
		userRoutes.GET("/me", userhttp.GetMeHandler(handler))
	}
}
