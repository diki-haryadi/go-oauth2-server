package oauthDto

import (
	oauthModel "github.com/diki-haryadi/go-micro-template/internal/oauth/domain/model"
	validator "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserRequestDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
	RoleID   string `json:"role_id"`
}

func (g *UserRequestDto) GetFieldsUser(ctx echo.Context) *UserRequestDto {
	return &UserRequestDto{
		Username: ctx.FormValue("username"),
		Password: ctx.FormValue("password"),
		RoleID:   ctx.FormValue("role_id"),
	}
}

func (g *UserRequestDto) ToModelUser(clientID uuid.UUID) *oauthModel.Client {
	return &oauthModel.Client{
		Common: oauthModel.Common{
			ID: clientID,
		},
	}
}

func (caDto *UserRequestDto) ValidateUserDto() error {
	return validator.ValidateStruct(caDto,
		validator.Field(
			&caDto.Username,
			validator.Required,
			validator.Length(5, 50),
		),
		validator.Field(
			&caDto.Password,
			validator.Required,
			validator.Length(6, 30),
		),
		validator.Field(
			&caDto.RoleID,
			validator.Required,
		),
	)
}

// UserResponse ...
type UserResponse struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
