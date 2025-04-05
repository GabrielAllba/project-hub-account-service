package response

import (
	"project-hub/account-service/internal/domain/entity"
	"project-hub/account-service/internal/dto/schema"
)

type RegisterResponse struct {
	ErrorSchema  schema.ErrorSchema `json:"errorSchema"`
	OutputSchema entity.User        `json:"outputSchema"`
}
