package user

import (
	"net/http"
	"strings"

	"project-hub/account-service/internal/domain/entity"
	userresponse "project-hub/account-service/internal/dto/response/user"
	"project-hub/account-service/internal/dto/schema"

	"github.com/gin-gonic/gin"
)

// LogoutHandler godoc
// @Summary      Logout user
// @Description  Invalidate the user's JWT by blacklisting the token
// @Tags         auth
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Success      200  {object}  userresponse.LogoutResponse  "Logout successful"
// @Failure      401  {object}  userresponse.LogoutResponse  "Unauthorized"
// @Failure      500  {object}  userresponse.LogoutResponse  "Internal server error"
// @Router       /auth/logout [post]
func LogoutHandler(h *UserHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, userresponse.LogoutResponse{
				ErrorSchema: schema.NewError("06", "Token tidak ditemukan", "Authorization token is missing"),
			})
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, userresponse.LogoutResponse{
				ErrorSchema: schema.NewError("06", "Format token tidak valid", "Invalid token format"),
			})
			return
		}

		tokenString := tokenParts[1]

		res, err := h.UserController.Logout(tokenString)
		if err != nil {
			var errorSchema schema.ErrorSchema

			switch err {
			case entity.ErrUnauthorized:
				errorSchema = schema.NewError("06", "Token tidak valid", "Invalid token")
			default:
				errorSchema = schema.NewError("99", "Kesalahan server", "Internal server error")
			}

			c.JSON(http.StatusInternalServerError, userresponse.LogoutResponse{
				ErrorSchema: errorSchema,
			})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}
