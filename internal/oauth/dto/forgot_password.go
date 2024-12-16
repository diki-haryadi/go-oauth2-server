package oauthDto

import (
	oauthModel "github.com/diki-haryadi/go-micro-template/internal/oauth/domain/model"
	validator "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ForgotPasswordRequest struct {
	UUID     string `json:"uuid"`
	Password string `json:"password"`
}

func (g *ForgotPasswordRequest) GetFieldsForgotPassword(ctx echo.Context) *ForgotPasswordRequest {
	return &ForgotPasswordRequest{
		UUID:     ctx.FormValue("uuid"),
		Password: ctx.FormValue("password"),
	}
}

func (g *ForgotPasswordRequest) ToModel(clientID uuid.UUID) *oauthModel.Client {
	return &oauthModel.Client{
		Common: oauthModel.Common{
			ID: clientID,
		},
	}
}

func (caDto *ForgotPasswordRequest) ValidateForgotPasswordDto() error {
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
	)
}

type ForgotPasswordResponse struct {
	Status bool `json:"status"`
}
