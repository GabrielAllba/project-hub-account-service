package response

import (
	"project-hub/account-service/internal/dto/output_schema"
	"project-hub/account-service/internal/dto/schema"
)

type LoginResponse struct {
	ErrorSchema  schema.ErrorSchema           `json:"errorSchema"`
	OutputSchema output_schema.UserLoginOutputSchema `json:"outputSchema"`
}

