package oauthDto

import (
	oauthModel "github.com/diki-haryadi/go-micro-template/internal/oauth/domain/model"
	validator "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UpdateUsernameRequest struct {
	UUID     string `json:"uuid"`
	Username string `json:"username"`
}

func (g *UpdateUsernameRequest) GetFieldsUpdateUsername(ctx echo.Context) *UpdateUsernameRequest {
	return &UpdateUsernameRequest{
		UUID:     ctx.FormValue("uuid"),
		Username: ctx.FormValue("username"),
	}
}

func (g *UpdateUsernameRequest) ToModel(userID uuid.UUID) *oauthModel.Users {
	return &oauthModel.Users{
		Common: oauthModel.Common{
			ID: userID,
		},
	}
}

func (caDto *UpdateUsernameRequest) ValidateUsernameDto() error {
	return validator.ValidateStruct(caDto,
		validator.Field(
			&caDto.UUID,
			validator.Required,
		),
		validator.Field(
			&caDto.Username,
			validator.Required,
			validator.Length(6, 30),
		),
	)
}

type UsernameResponse struct {
	Status bool `json:"status"`
}
