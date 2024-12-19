package authDto

import (
	oauthModel "github.com/diki-haryadi/go-micro-template/internal/authentication/domain/model"
	validator "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ChangePasswordRequest struct {
	UUID        string `json:"uuid"`
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
}

func (g *ChangePasswordRequest) GetFieldsChangePassword(ctx echo.Context) *ChangePasswordRequest {
	return &ChangePasswordRequest{
		UUID:        ctx.FormValue("uuid"),
		Password:    ctx.FormValue("password"),
		NewPassword: ctx.FormValue("new_password"),
	}
}

func (g *ChangePasswordRequest) GetFieldsChangePasswordValue(uuid, password, new_password string) *ChangePasswordRequest {
	return &ChangePasswordRequest{
		UUID:        uuid,
		Password:    password,
		NewPassword: new_password,
	}
}

func (g *ChangePasswordRequest) ToModel(clientID uuid.UUID) *oauthModel.Client {
	return &oauthModel.Client{
		Common: oauthModel.Common{
			ID: clientID,
		},
	}
}

func (caDto *ChangePasswordRequest) ValidateChangePasswordDto() error {
	return validator.ValidateStruct(caDto,
		validator.Field(
			&caDto.UUID,
			validator.Required,
		),
		validator.Field(
			&caDto.Password,
			validator.Required,
			validator.Length(6, 30),
		),
		validator.Field(
			&caDto.NewPassword,
			validator.Required,
			validator.Length(6, 30),
		),
	)
}

type ChangePasswordResponse struct {
	Status bool `json:"status"`
}
