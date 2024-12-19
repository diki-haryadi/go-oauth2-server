package oauthDto

import (
	"fmt"
	"github.com/diki-haryadi/go-micro-template/config"
	oauthModel "github.com/diki-haryadi/go-micro-template/internal/oauth/domain/model"
	"github.com/diki-haryadi/go-micro-template/pkg"
	validator "github.com/go-ozzo/ozzo-validation"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"time"
)

type RefreshTokenRequestDto struct {
	GrantType    string `json:"grant_type"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

func (g *RefreshTokenRequestDto) GetFields(ctx echo.Context) *RefreshTokenRequestDto {
	return &RefreshTokenRequestDto{
		GrantType:    ctx.FormValue("grant_type"),
		RefreshToken: ctx.FormValue("refresh_token"),
		Scope:        ctx.FormValue("scope"),
	}
}

func (g *RefreshTokenRequestDto) GetFieldsValue(refreshToken string, scope string) *RefreshTokenRequestDto {
	return &RefreshTokenRequestDto{
		GrantType:    "-",
		RefreshToken: refreshToken,
		Scope:        scope,
	}
}

func (g *RefreshTokenRequestDto) ToModel(clientID uuid.UUID) *oauthModel.Client {
	return &oauthModel.Client{
		Common: oauthModel.Common{
			ID: clientID,
		},
	}
}

func (caDto *RefreshTokenRequestDto) ValidateRefreshTokenDto() error {
	return validator.ValidateStruct(caDto,
		validator.Field(
			&caDto.GrantType,
			validator.Required,
		),
		validator.Field(
			&caDto.RefreshToken,
			validator.Required,
		),
		validator.Field(
			&caDto.Scope,
			validator.Required,
		),
	)
}

func NewOauthRefreshToken(client *oauthModel.Client, user *oauthModel.Users, expiresIn int, scope string) (*oauthModel.RefreshToken, error) {
	tokenID := uuid.New()
	claims := TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        tokenID.String(),
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(time.Duration(expiresIn) * time.Second)),
		},
		ClientID:  fmt.Sprint(client.ID),
		Scope:     scope,
		TokenType: "refresh_token",
	}

	if user != nil {
		claims.UserID = fmt.Sprint(user.ID)
	}

	token, err := generateJWTToken(claims, config.BaseConfig.App.ConfigOauth.JWTSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to generate refresh token: %w", err)
	}

	refreshToken := &oauthModel.RefreshToken{
		Common: oauthModel.Common{
			ID:        tokenID,
			CreatedAt: time.Now().UTC(),
		},
		ClientID:  pkg.StringOrNull(fmt.Sprint(client.ID)),
		Token:     token,
		ExpiresAt: time.Now().UTC().Add(time.Duration(expiresIn) * time.Second),
		Scope:     scope,
	}

	if user != nil {
		refreshToken.UserID = pkg.StringOrNull(fmt.Sprint(user.ID))
	}

	return refreshToken, nil
}
