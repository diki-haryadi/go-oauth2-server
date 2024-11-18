package oauthUseCase

import (
	"context"
	"errors"
	"github.com/diki-haryadi/go-micro-template/config"
	oauthDomain "github.com/diki-haryadi/go-micro-template/internal/oauth/domain/model"
	oauthDto "github.com/diki-haryadi/go-micro-template/internal/oauth/dto"
	"github.com/diki-haryadi/go-micro-template/pkg"
)

var (
	// ErrInvalidUsernameOrPassword ...
	ErrInvalidUsernameOrPassword = errors.New("Invalid username or password")
)

func (uc *useCase) PasswordGrant(ctx context.Context, username, password string, scope string, client *oauthDomain.Client) (*oauthDto.AccessTokenResponse, error) {
	// Get the scope string
	scope, err := uc.GetScope(ctx, scope)
	if err != nil {
		return nil, err
	}

	// Authenticate the user
	user, err := uc.AuthUser(username, password)
	if err != nil {
		// For security reasons, return a general error message
		return nil, ErrInvalidUsernameOrPassword
	}

	// Log in the user
	accessToken, refreshToken, err := uc.Login(ctx, client, user, scope)
	if err != nil {
		return nil, err
	}

	// Create response
	accessTokenResponse, err := oauthDto.NewAccessTokenResponse(
		accessToken,
		refreshToken,
		config.BaseConfig.App.ConfigOauth.Oauth.AccessTokenLifetime,
		pkg.Bearer,
	)
	if err != nil {
		return nil, err
	}

	return accessTokenResponse, nil
}
