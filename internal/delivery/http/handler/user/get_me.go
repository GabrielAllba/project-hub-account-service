package user

import (
	"net/http"

	"project-hub/account-service/internal/common/contextutil"
	"project-hub/account-service/internal/domain/entity"

	userresponse "project-hub/account-service/internal/dto/response/user"
	"project-hub/account-service/internal/dto/schema"

	"github.com/gin-gonic/gin"
)

// GetMeHandler godoc
// @Summary      Get current user
// @Description  Get currently authenticated user's info
// @Tags         users
// @Security     BearerAuth
// @Produce      json
// @Success      200  {object}  userresponse.GetMeResponse  "Success"
// @Failure      401  {object}  userresponse.GetMeResponse  "Unauthorized"
// @Failure      404  {object}  userresponse.GetMeResponse  "User not found"
// @Router       /users/me [get]
func GetMeHandler(h *UserHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, ok := contextutil.GetUserIDFromContext(c)
		if !ok {
			c.JSON(http.StatusUnauthorized, userresponse.GetMeResponse{
				ErrorSchema:  schema.NewError("06", "Tidak diotorisasi", "Unauthorized"),
				OutputSchema: entity.User{},
			})
			return
		}

		res, err := h.UserController.GetUser(userID)
		if err != nil {
			switch err {
			case entity.ErrNotFound:
				c.JSON(http.StatusOK, userresponse.GetMeResponse{
					ErrorSchema:  schema.NewError("07", "Pengguna tidak ditemukan", "User not found"),
					OutputSchema: entity.User{},
				})
			default:
				c.JSON(http.StatusOK, userresponse.GetMeResponse{
					ErrorSchema:  schema.NewError("05", "Gagal mengambil data pengguna", "Error fetching user"),
					OutputSchema: entity.User{},
				})
			}
			return
		}

		c.JSON(http.StatusOK, userresponse.GetMeResponse{
			ErrorSchema:  schema.NewSuccess(),
			OutputSchema: res.OutputSchema,
		})
	}
}
