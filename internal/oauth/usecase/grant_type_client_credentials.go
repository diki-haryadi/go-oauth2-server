package oauthUseCase

import (
	"context"
	"github.com/diki-haryadi/go-micro-template/config"
	oauthDomain "github.com/diki-haryadi/go-micro-template/internal/oauth/domain/model"
	oauthDto "github.com/diki-haryadi/go-micro-template/internal/oauth/dto"
	"github.com/diki-haryadi/go-micro-template/pkg"
)

func (uc *useCase) ClientCredentialsGrant(ctx context.Context, scope string, client *oauthDomain.Client) (*oauthDto.AccessTokenResponse, error) {
	// Get the scope string
	scope, err := uc.GetScope(ctx, scope)
	if err != nil {
		return nil, err
	}

	// Create a new access token
	accessToken, err := uc.repository.GrantAccessToken(
		client,
		nil, // empty user
		config.BaseConfig.App.ConfigOauth.Oauth.AccessTokenLifetime, // expires in
		scope,
	)
	if err != nil {
		return nil, err
	}

	// Create response
	accessTokenResponse, err := oauthDto.NewAccessTokenResponse(
		accessToken,
		nil, // refresh token
		config.BaseConfig.App.ConfigOauth.Oauth.AccessTokenLifetime,
		pkg.Bearer,
	)
	if err != nil {
		return nil, err
	}

	return accessTokenResponse, nil
}
