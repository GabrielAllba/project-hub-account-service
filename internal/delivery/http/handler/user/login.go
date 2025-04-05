package user

import (
	"net/http"

	"project-hub/account-service/internal/domain/entity"
	"project-hub/account-service/internal/dto/output_schema"
	"project-hub/account-service/internal/dto/request"
	userresponse "project-hub/account-service/internal/dto/response/user"
	"project-hub/account-service/internal/dto/schema"

	"github.com/gin-gonic/gin"
)

// LoginHandler godoc
// @Summary      Login user
// @Description  Authenticate user and return JWT
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        credentials  body      request.LoginRequest       true  "Login credentials"
// @Success      200          {object}  userresponse.LoginResponse  "Login successful"
// @Failure      400          {object}  userresponse.LoginResponse  "Bad request"
// @Failure      401          {object}  userresponse.LoginResponse  "Unauthorized"
// @Failure      500          {object}  userresponse.LoginResponse  "Internal server error"
// @Router       /auth/login [post]
func LoginHandler(h *UserHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request.LoginRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, userresponse.LoginResponse{
				ErrorSchema:  schema.NewError("01", "Permintaan tidak valid", "Invalid request payload"),
				OutputSchema: output_schema.UserLoginOutputSchema{},
			})
			return
		}

		if err := h.Validator.Struct(req); err != nil {
			c.JSON(http.StatusBadRequest, userresponse.LoginResponse{
				ErrorSchema:  schema.NewError("02", "Validasi gagal", err.Error()),
				OutputSchema: output_schema.UserLoginOutputSchema{},
			})
			return
		}

		res, err := h.UserController.Login(&req)
		if err != nil {
			var errorSchema schema.ErrorSchema

			switch err {
			case entity.ErrInvalidInput:
				errorSchema = schema.NewError("03", "Input tidak valid", err.Error())
			case entity.ErrUnauthorized:
				errorSchema = schema.NewError("06", "Email atau password salah", "Invalid credentials")
			default:
				errorSchema = schema.NewError("05", "Gagal login", "Error during login")
			}

			c.JSON(http.StatusOK, userresponse.LoginResponse{
				ErrorSchema:  errorSchema,
				OutputSchema: output_schema.UserLoginOutputSchema{},
			})
			return
		}

		c.JSON(http.StatusOK, userresponse.LoginResponse{
			ErrorSchema:  schema.NewSuccess(),
			OutputSchema: res.OutputSchema,
		})
	}
}
