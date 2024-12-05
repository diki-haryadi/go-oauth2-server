package oauthDto

import (
	"fmt"
	"github.com/diki-haryadi/go-micro-template/config"
	oauthModel "github.com/diki-haryadi/go-micro-template/internal/oauth/domain/model"
	"github.com/diki-haryadi/go-micro-template/pkg"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

// AccessTokenResponse ...
type AccessTokenResponse struct {
	UserID       string `json:"user_id,omitempty"`
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

// NewAccessTokenResponse ...
func NewAccessTokenResponse(accessToken *oauthModel.AccessToken, refreshToken *oauthModel.RefreshToken, lifetime int, theTokenType string) (*AccessTokenResponse, error) {
	response := &AccessTokenResponse{
		AccessToken: accessToken.Token,
		ExpiresIn:   lifetime,
		TokenType:   theTokenType,
		Scope:       accessToken.Scope,
	}
	if accessToken.UserID.Valid {
		response.UserID = accessToken.UserID.String
	}
	if refreshToken != nil {
		response.RefreshToken = refreshToken.Token
	}
	return response, nil
}

func NewOauthAccessToken(client *oauthModel.Client, user *oauthModel.Users, expiresIn int, scope string) (*oauthModel.AccessToken, error) {
	tokenID := uuid.New()
	claims := TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        tokenID.String(),
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(time.Duration(expiresIn) * time.Second)),
		},
		ClientID:  fmt.Sprint(client.ID),
		Scope:     scope,
		TokenType: "access_token",
	}

	if user != nil {
		claims.UserID = fmt.Sprint(user.ID)
	}

	token, err := generateJWTToken(claims, config.BaseConfig.App.ConfigOauth.JWTSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %w", err)
	}

	accessToken := &oauthModel.AccessToken{
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
		accessToken.UserID = pkg.StringOrNull(fmt.Sprint(user.ID))
	}

	return accessToken, nil
}
