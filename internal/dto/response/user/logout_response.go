package response

import (
	"project-hub/account-service/internal/dto/schema"
)

type LogoutResponse struct {
	ErrorSchema schema.ErrorSchema `json:"errorSchema"`
}
