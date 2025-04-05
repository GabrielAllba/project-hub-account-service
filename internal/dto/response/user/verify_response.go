package response

import (
	"project-hub/account-service/internal/domain/entity"
	"project-hub/account-service/internal/dto/schema"
)

type VerifyUserResponse struct {
	ErrorSchema  schema.ErrorSchema `json:"error_schema"`
	OutputSchema entity.User        `json:"outputSchema"`
}
