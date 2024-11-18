package oauthUseCase

import (
	"context"
	"errors"
	oauthDomain "github.com/diki-haryadi/go-micro-template/internal/oauth/domain/model"
	oauthDto "github.com/diki-haryadi/go-micro-template/internal/oauth/dto"
	"github.com/diki-haryadi/go-micro-template/pkg"
)

const (
	// AccessTokenHint ...
	AccessTokenHint = "access_token"
	// RefreshTokenHint ...
	RefreshTokenHint = "refresh_token"
)

var (
	// ErrTokenHintInvalid ...
	ErrTokenHintInvalid = errors.New("Invalid token hint")
)

func (uc *useCase) IntrospectToken(ctx context.Context, token, tokenTypeHint string, client *oauthDomain.Client) (*oauthDto.IntrospectResponse, error) {
	// Default to access token hint
	if tokenTypeHint == "" {
		tokenTypeHint = AccessTokenHint
	}

	switch tokenTypeHint {
	case AccessTokenHint:
		accessToken, err := uc.repository.Authenticate(token)
		if err != nil {
			return nil, err
		}
		return uc.NewIntrospectResponseFromAccessToken(accessToken)
	case RefreshTokenHint:
		refreshToken, err := uc.repository.GetValidRefreshToken(token, client)
		if err != nil {
			return nil, err
		}
		return uc.NewIntrospectResponseFromRefreshToken(refreshToken)
	default:
		return nil, ErrTokenHintInvalid
	}
}

// NewIntrospectResponseFromAccessToken creates an introspect response from an access token
func (uc *useCase) NewIntrospectResponseFromAccessToken(accessToken *oauthDomain.AccessToken) (*oauthDto.IntrospectResponse, error) {
	var introspectResponse = &oauthDto.IntrospectResponse{
		Active:    true,
		Scope:     accessToken.Scope,
		TokenType: pkg.Bearer,
		ExpiresAt: int(accessToken.ExpiresAt.Unix()),
	}

	// Fetch the client using the FetchClientByClientID method
	if accessToken.ClientID.Valid {
		client, err := uc.repository.FetchClientByClientID(accessToken.ClientID.String)
		if err != nil {
			return nil, err
		}
		introspectResponse.ClientID = client.Key
	}

	// Fetch the user using the FetchUserByUserID method
	if accessToken.UserID.Valid {
		user, err := uc.repository.FetchUserByUserID(accessToken.UserID.String)
		if err != nil {
			return nil, err
		}
		introspectResponse.Username = user.Username
	}

	return introspectResponse, nil
}

// NewIntrospectResponseFromRefreshToken creates an introspect response from a refresh token
func (uc *useCase) NewIntrospectResponseFromRefreshToken(refreshToken *oauthDomain.RefreshToken) (*oauthDto.IntrospectResponse, error) {
	var introspectResponse = &oauthDto.IntrospectResponse{
		Active:    true,
		Scope:     refreshToken.Scope,
		TokenType: pkg.Bearer,
		ExpiresAt: int(refreshToken.ExpiresAt.Unix()),
	}

	// Fetch the client using the FetchClientByClientID method
	if refreshToken.ClientID.Valid {
		client, err := uc.repository.FetchClientByClientID(refreshToken.ClientID.String)
		if err != nil {
			return nil, err
		}
		introspectResponse.ClientID = client.Key
	}

	// Fetch the user using the FetchUserByUserID method
	if refreshToken.UserID.Valid {
		user, err := uc.repository.FetchUserByUserID(refreshToken.UserID.String)
		if err != nil {
			return nil, err
		}
		introspectResponse.Username = user.Username
	}

	return introspectResponse, nil
}
