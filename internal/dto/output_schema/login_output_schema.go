package output_schema

import "project-hub/account-service/internal/domain/entity"

type UserLoginOutputSchema struct {
	Token string      `json:"token"`
	User  entity.User `json:"user"`
}
