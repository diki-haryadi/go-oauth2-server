package oauthGrpcController

import (
	"context"
	"github.com/diki-haryadi/go-micro-template/config"
	oauthDomain "github.com/diki-haryadi/go-micro-template/internal/oauth/domain"
	oauthModel "github.com/diki-haryadi/go-micro-template/internal/oauth/domain/model"
	oauthDto "github.com/diki-haryadi/go-micro-template/internal/oauth/dto"
	"github.com/diki-haryadi/go-micro-template/pkg/response"
	oauth2 "github.com/diki-haryadi/protobuf-ecomerce/oauth2_server_service/oauth2/v1"
)

type controller struct {
	useCase oauthDomain.UseCase
}

func NewController(uc oauthDomain.UseCase) oauthDomain.GrpcController {
	return &controller{
		useCase: uc,
	}
}

func (c *controller) BasicAuthClient(ctx context.Context, clientID, secret string) (*oauthModel.Client, error) {

	// Authenticate the client
	client, err := c.useCase.AuthClient(ctx, clientID, secret)
	if err != nil {
		return nil, response.ErrInvalidClientIDOrSecret
	}

	return client, nil
}

func (c *controller) PasswordGrant(ctx context.Context, req *oauth2.PasswordGrantRequest) (*oauth2.PasswordGrantResponse, error) {
	aDto := new(oauthDto.PasswordGrantRequestDto).GetFieldsValue(req.Username, req.Password, req.Scope)
	if err := aDto.ValidatePasswordDto(); err != nil {
		return &oauth2.PasswordGrantResponse{}, err
	}

	client, err := c.BasicAuthClient(ctx, req.ClientId, req.ClientSecret)
	if err != nil {
		return &oauth2.PasswordGrantResponse{}, err
	}

	acGrant, err := c.useCase.PasswordGrant(
		ctx,
		aDto.Username,
		aDto.Password,
		aDto.Scope,
		aDto.ToModel(client.ID))

	if err != nil {
		return &oauth2.PasswordGrantResponse{}, err
	}

	return &oauth2.PasswordGrantResponse{
		AccessToken:  acGrant.AccessToken,
		TokenType:    acGrant.TokenType,
		ExpiresIn:    int32(acGrant.ExpiresIn),
		Scope:        acGrant.Scope,
		RefreshToken: acGrant.RefreshToken,
	}, nil
}

func (c *controller) AuthorizationCodeGrant(ctx context.Context, req *oauth2.AuthorizationCodeGrantRequest) (*oauth2.AuthorizationCodeGrantResponse, error) {
	aDto := new(oauthDto.AuthorizationCodeGrantRequestDto).GetFieldsValue(req.Code, req.RedirectUri, req.ClientId)
	if err := aDto.ValidateAuthorizationCodeDto(); err != nil {
		return &oauth2.AuthorizationCodeGrantResponse{}, err
	}

	client, err := c.BasicAuthClient(ctx, req.ClientId, req.ClientSecret)
	if err != nil {
		return &oauth2.AuthorizationCodeGrantResponse{}, err
	}

	acGrant, err := c.useCase.AuthorizationCodeGrant(
		ctx,
		aDto.Code,
		aDto.RedirectUri,
		aDto.ToModel(client.ID))

	if err != nil {
		return &oauth2.AuthorizationCodeGrantResponse{}, err
	}
	return &oauth2.AuthorizationCodeGrantResponse{
		AccessToken:  acGrant.AccessToken,
		TokenType:    acGrant.TokenType,
		ExpiresIn:    int32(acGrant.ExpiresIn),
		Scope:        acGrant.Scope,
		RefreshToken: acGrant.RefreshToken,
	}, nil
}

func (c *controller) ClientCredentialsGrant(ctx context.Context, req *oauth2.ClientCredentialsGrantRequest) (*oauth2.ClientCredentialsGrantResponse, error) {
	aDto := new(oauthDto.ClientCredentialsGrantRequestDto).GetFieldsValue(req.Scope)
	if err := aDto.ValidateClientCredentialsDto(); err != nil {
		return &oauth2.ClientCredentialsGrantResponse{}, err
	}

	client, err := c.BasicAuthClient(ctx, req.ClientId, req.ClientSecret)
	if err != nil {
		return &oauth2.ClientCredentialsGrantResponse{}, err
	}

	acGrant, err := c.useCase.GrantAccessToken(
		ctx,
		client,
		nil,
		config.BaseConfig.App.ConfigOauth.Oauth.AccessTokenLifetime,
		aDto.Scope)

	if err != nil {
		return &oauth2.ClientCredentialsGrantResponse{}, err
	}

	return &oauth2.ClientCredentialsGrantResponse{
		AccessToken:  acGrant.AccessToken,
		TokenType:    acGrant.TokenType,
		ExpiresIn:    int32(acGrant.ExpiresIn),
		Scope:        acGrant.Scope,
		RefreshToken: acGrant.RefreshToken,
	}, err
}

func (c *controller) RefreshTokenGrant(ctx context.Context, req *oauth2.RefreshTokenGrantRequest) (*oauth2.RefreshTokenGrantResponse, error) {
	aDto := new(oauthDto.RefreshTokenRequestDto).GetFieldsValue(req.RefreshToken, req.Scope)
	if err := aDto.ValidateRefreshTokenDto(); err != nil {
		return &oauth2.RefreshTokenGrantResponse{}, nil
	}

	client, err := c.BasicAuthClient(ctx, req.ClientId, req.ClientSecret)
	if err != nil {
		return &oauth2.RefreshTokenGrantResponse{}, nil
	}

	acGrant, err := c.useCase.ClientCredentialsGrant(
		ctx,
		aDto.RefreshToken,
		aDto.Scope,
		aDto.ToModel(client.ID))

	if err != nil {
		return &oauth2.RefreshTokenGrantResponse{}, nil
	}
	return &oauth2.RefreshTokenGrantResponse{
		AccessToken:  acGrant.AccessToken,
		TokenType:    acGrant.TokenType,
		ExpiresIn:    int32(acGrant.ExpiresIn),
		Scope:        acGrant.Scope,
		RefreshToken: acGrant.RefreshToken,
	}, nil
}
