package user

import (
	"net/http"

	"project-hub/account-service/internal/domain/entity"
	userresponse "project-hub/account-service/internal/dto/response/user"
	"project-hub/account-service/internal/dto/schema"

	"github.com/gin-gonic/gin"
)

// VerifyUserHandler godoc
// @Summary      Verify a user
// @Description  Mark a user as verified
// @Tags         users
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Success      200  {object}  userresponse.VerifyUserResponse  "User verified successfully"
// @Failure      401  {object}  userresponse.VerifyUserResponse  "Unauthorized"
// @Failure      500  {object}  userresponse.VerifyUserResponse  "Internal server error"
// @Router       /users/verify [put]
func VerifyUserHandler(h *UserHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, userresponse.VerifyUserResponse{
				ErrorSchema: schema.NewError("06", "Unauthorized", "Unauthorized"),
			})
			return
		}

		userResponse, err := h.UserController.VerifyUser(userID.(uint))
		if err != nil {
			c.JSON(http.StatusInternalServerError, userresponse.VerifyUserResponse{
				ErrorSchema: schema.NewError("99", "Failed to verify user", "Internal server error"),
			})
			return
		}

		userEntity := entity.User{
			ID:         userResponse.OutputSchema.ID,
			Email:      userResponse.OutputSchema.Email,
			VerifiedAt: userResponse.OutputSchema.VerifiedAt,
			CreatedAt:  userResponse.OutputSchema.CreatedAt,
			UpdatedAt:  userResponse.OutputSchema.UpdatedAt,
		}

		c.JSON(http.StatusOK, userresponse.VerifyUserResponse{
			ErrorSchema:  schema.NewError("00", "Verification successful", "Verification successful"),
			OutputSchema: userEntity,
		})
	}
}
