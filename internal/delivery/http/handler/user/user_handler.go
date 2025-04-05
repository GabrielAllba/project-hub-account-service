package user

import (
	"project-hub/account-service/internal/controller"

	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	UserController controller.UserController
	Validator      *validator.Validate
}
