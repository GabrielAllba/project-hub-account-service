package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

	"project-hub/account-service/internal/domain/entity"
	userresponse "project-hub/account-service/internal/dto/response/user"
	"project-hub/account-service/internal/dto/schema"
	"project-hub/account-service/pkg/config"
)

func AuthMiddleware(cfg *config.Config, userRepo entity.UserRepository, blacklistedTokenRepo entity.BlacklistedTokenRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, userresponse.GetMeResponse{
				ErrorSchema:  schema.NewError("06", "Token tidak ditemukan", "Authorization token is missing"),
				OutputSchema: entity.User{},
			})
			c.Abort()
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, userresponse.GetMeResponse{
				ErrorSchema:  schema.NewError("06", "Format token tidak valid", "Invalid token format"),
				OutputSchema: entity.User{},
			})
			c.Abort()
			return
		}

		tokenString := tokenParts[1]

		isBlacklisted, err := blacklistedTokenRepo.IsTokenBlacklisted(tokenString)
		if err != nil {
			c.JSON(http.StatusInternalServerError, userresponse.GetMeResponse{
				ErrorSchema:  schema.NewError("99", "Kesalahan server", "Internal server error"),
				OutputSchema: entity.User{},
			})
			c.Abort()
			return
		}
		if isBlacklisted {
			c.JSON(http.StatusUnauthorized, userresponse.GetMeResponse{
				ErrorSchema:  schema.NewError("06", "Token telah diblacklist", "Token has been blacklisted"),
				OutputSchema: entity.User{},
			})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, entity.ErrUnauthorized
			}
			return []byte(cfg.JWTSecret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, userresponse.GetMeResponse{
				ErrorSchema:  schema.NewError("06", "Token tidak valid", "Invalid or expired token"),
				OutputSchema: entity.User{},
			})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, userresponse.GetMeResponse{
				ErrorSchema:  schema.NewError("06", "Klaim token tidak valid", "Invalid token claims"),
				OutputSchema: entity.User{},
			})
			c.Abort()
			return
		}

		userID, ok := claims["user_id"].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, userresponse.GetMeResponse{
				ErrorSchema:  schema.NewError("06", "User ID tidak ditemukan dalam token", "User ID not found in token"),
				OutputSchema: entity.User{},
			})
			c.Abort()
			return
		}

		c.Set("userID", uint(userID))
		c.Next()
	}
}
