package authDto

import (
	oauthModel "github.com/diki-haryadi/go-micro-template/internal/authentication/domain/model"
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

func (g *UserRequestDto) GetFieldsUserValue(username, password, role_id string) *UserRequestDto {
	return &UserRequestDto{
		Username: username,
		Password: password,
		RoleID:   role_id,
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
	UUID     string `json:"uuid"`
	Username string `json:"username"`
	Password string `json:"password"`
	RoleID   string `json:"role_id"`
}
