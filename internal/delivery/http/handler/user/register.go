package user

import (
	"net/http"

	"project-hub/account-service/internal/domain/entity"
	"project-hub/account-service/internal/dto/request"
	userresponse "project-hub/account-service/internal/dto/response/user"
	"project-hub/account-service/internal/dto/schema"

	"github.com/gin-gonic/gin"
)

// RegisterHandler godoc
// @Summary      Register a new user
// @Description  Create a new user account
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body      request.RegisterRequest       true  "User Registration Info"
// @Success      201   {object}  userresponse.RegisterResponse  "User created successfully"
// @Failure      400   {object}  userresponse.RegisterResponse  "Invalid input"
// @Failure      409   {object}  userresponse.RegisterResponse  "User already exists"
// @Failure      500   {object}  userresponse.RegisterResponse  "Internal server error"
// @Router       /auth/register [post]
func RegisterHandler(h *UserHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request.RegisterRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, userresponse.RegisterResponse{
				ErrorSchema:  schema.NewError("01", "Permintaan tidak valid", "Invalid request payload"),
				OutputSchema: entity.User{},
			})
			return
		}

		if err := h.Validator.Struct(req); err != nil {
			c.JSON(http.StatusBadRequest, userresponse.RegisterResponse{
				ErrorSchema:  schema.NewError("02", "Validasi gagal", err.Error()),
				OutputSchema: entity.User{},
			})
			return
		}

		res, err := h.UserController.Register(&req)
		if err != nil {

			var errorSchema schema.ErrorSchema

			switch err {
			case entity.ErrInvalidInput:
				errorSchema = schema.NewError("03", "Input tidak valid", err.Error())
			case entity.ErrAlreadyExists:
				errorSchema = schema.NewError("04", "Email sudah digunakan", "User with this email already exists")
			default:
				errorSchema = schema.NewError("05", "Gagal membuat pengguna", "Error creating user")
			}
			c.JSON(http.StatusOK, userresponse.RegisterResponse{
				ErrorSchema:  errorSchema,
				OutputSchema: entity.User{},
			})
			return
		}

		c.JSON(http.StatusCreated, userresponse.RegisterResponse{
			ErrorSchema:  schema.NewSuccess(),
			OutputSchema: res.OutputSchema,
		})
	}
}
