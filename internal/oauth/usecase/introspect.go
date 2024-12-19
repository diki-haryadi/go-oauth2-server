package oauthUseCase

import (
	"context"
	"github.com/diki-haryadi/go-micro-template/config"
	oauthDomain "github.com/diki-haryadi/go-micro-template/internal/oauth/domain/model"
	oauthDto "github.com/diki-haryadi/go-micro-template/internal/oauth/dto"
	"github.com/diki-haryadi/go-micro-template/pkg"
	"github.com/diki-haryadi/go-micro-template/pkg/constant"
	"github.com/diki-haryadi/go-micro-template/pkg/response"
	"time"
)

func (uc *useCase) IntrospectToken(ctx context.Context, token, tokenTypeHint string, client *oauthDomain.Client) (*oauthDto.IntrospectResponse, error) {
	if tokenTypeHint == "" {
		tokenTypeHint = constant.AccessTokenHint
	}

	claims, err := oauthDto.ValidateToken(token, config.BaseConfig.App.ConfigOauth.JWTSecret)
	if err != nil {
		return &oauthDto.IntrospectResponse{Active: false}, nil
	}

	if time.Now().After(claims.ExpiresAt.Time) {
		return &oauthDto.IntrospectResponse{Active: false}, nil
	}

	switch tokenTypeHint {
	case constant.AccessTokenHint:
		if claims.TokenType != "access_token" {
			return &oauthDto.IntrospectResponse{Active: false}, nil
		}
		accessToken, err := uc.repository.Authenticate(ctx, token)
		if err != nil {
			return &oauthDto.IntrospectResponse{Active: false}, nil
		}
		return uc.NewIntrospectResponseFromAccessToken(ctx, accessToken, claims)

	case constant.RefreshTokenHint:
		if claims.TokenType != "refresh_token" {
			return &oauthDto.IntrospectResponse{Active: false}, nil
		}
		refreshToken, err := uc.repository.GetValidRefreshToken(ctx, token, client)
		if err != nil {
			return &oauthDto.IntrospectResponse{Active: false}, nil
		}
		return uc.NewIntrospectResponseFromRefreshToken(ctx, refreshToken, claims)

	default:
		return nil, response.ErrTokenHintInvalid
	}
}

func (uc *useCase) NewIntrospectResponseFromAccessToken(ctx context.Context, accessToken *oauthDomain.AccessToken, claims *oauthDto.TokenClaims) (*oauthDto.IntrospectResponse, error) {
	introspectResponse := &oauthDto.IntrospectResponse{
		Active:    true,
		Scope:     claims.Scope,
		TokenType: pkg.Bearer,
		ExpiresAt: int(claims.ExpiresAt.Unix()),
		IssuedAt:  int(claims.IssuedAt.Unix()),
		JTI:       claims.ID,
	}

	if claims.ClientID != "" {
		client, err := uc.repository.FetchClientByClientID(ctx, claims.ClientID)
		if err != nil {
			return nil, err
		}
		introspectResponse.ClientID = client.Key
	}

	if claims.UserID != "" {
		user, err := uc.repository.FetchUserByUserID(ctx, claims.UserID)
		if err != nil {
			return nil, err
		}
		introspectResponse.Username = user.Username
		introspectResponse.Sub = claims.UserID
	}

	return introspectResponse, nil
}

func (uc *useCase) NewIntrospectResponseFromRefreshToken(ctx context.Context, refreshToken *oauthDomain.RefreshToken, claims *oauthDto.TokenClaims) (*oauthDto.IntrospectResponse, error) {
	introspectResponse := &oauthDto.IntrospectResponse{
		Active:    true,
		Scope:     claims.Scope,
		TokenType: "refresh_token",
		ExpiresAt: int(claims.ExpiresAt.Unix()),
		IssuedAt:  int(claims.IssuedAt.Unix()),
		JTI:       claims.ID,
	}

	if claims.ClientID != "" {
		client, err := uc.repository.FetchClientByClientID(ctx, claims.ClientID)
		if err != nil {
			return nil, err
		}
		introspectResponse.ClientID = client.Key
	}

	if claims.UserID != "" {
		user, err := uc.repository.FetchUserByUserID(ctx, claims.UserID)
		if err != nil {
			return nil, err
		}
		introspectResponse.Username = user.Username
		introspectResponse.Sub = claims.UserID
	}

	return introspectResponse, nil
}
