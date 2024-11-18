package oauthUseCase

import (
	"context"
	"github.com/diki-haryadi/go-micro-template/config"
	oauthDomain "github.com/diki-haryadi/go-micro-template/internal/oauth/domain/model"
)

// Login creates an access token and refresh token for a user (logs him/her in)
func (uc *useCase) Login(ctx context.Context, client *oauthDomain.Client, user *oauthDomain.Users, scope string) (*oauthDomain.AccessToken, *oauthDomain.RefreshToken, error) {
	// Return error if user's role is not allowed to use this service
	if !uc.IsRoleAllowed(user.RoleID.String) {
		// For security reasons, return a general error message
		return nil, nil, ErrInvalidUsernameOrPassword
	}

	// Create a new access token
	accessToken, err := uc.repository.GrantAccessToken(
		client,
		user,
		config.BaseConfig.App.ConfigOauth.Oauth.AccessTokenLifetime, // expires in
		scope,
	)
	if err != nil {
		return nil, nil, err
	}

	// Create or retrieve a refresh token
	refreshToken, err := uc.repository.GetOrCreateRefreshToken(
		client,
		user,
		config.BaseConfig.App.ConfigOauth.Oauth.AccessTokenLifetime, // expires in
		scope,
	)
	if err != nil {
		return nil, nil, err
	}

	return accessToken, refreshToken, nil
}
