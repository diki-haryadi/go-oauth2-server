package oauthHttpController

import (
	"errors"
	oauthModel "github.com/diki-haryadi/go-micro-template/internal/oauth/domain/model"
	"net/http"

	"github.com/labstack/echo/v4"

	oauthDomain "github.com/diki-haryadi/go-micro-template/internal/oauth/domain"
	oauthDto "github.com/diki-haryadi/go-micro-template/internal/oauth/dto"
	oauthException "github.com/diki-haryadi/go-micro-template/internal/oauth/exception"
)

type controller struct {
	useCase oauthDomain.UseCase
}

func NewController(uc oauthDomain.UseCase) oauthDomain.HttpController {
	return &controller{
		useCase: uc,
	}
}

func (c controller) Tokens(ctx echo.Context) error {
	if err := ctx.Request().ParseForm(); err != nil {
		//response.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	grantTypes := map[string]func(ctx echo.Context) error{
		"authorization_code": c.AuthorizationCodeGrant,
		"password":           c.PasswordGrant,
		"client_credentials": c.ClientCredentialsGrant,
		"refresh_token":      c.RefreshTokenGrant,
	}

	// Check the grant type
	grantHandler, ok := grantTypes[ctx.Request().Form.Get("grant_type")]
	if !ok {
		//response.Error(w, ErrInvalidGrantType.Error(), http.StatusBadRequest)
		return errors.New("invalid grant type")
	}

	// Grant processing
	err := grantHandler(ctx)
	if err != nil {
		//response.Error(w, err.Error(), getErrStatusCode(err))
		return err
	}

	return ctx.JSON(http.StatusOK, nil)
}

func (c controller) AuthorizationCodeGrant(ctx echo.Context) error {
	aDto := new(oauthDto.GrantAuthorizationCodeRequestDto)
	if err := ctx.Bind(aDto); err != nil {
		return oauthException.AuthorizationCodeGrantBindingExc()
	}

	if err := aDto.ValidateAuthorizationCodeDto(); err != nil {
		return oauthException.AuthorizationCodeGrantValidationExc(err)
	}

	client, err := c.BasicAuthClient(ctx)
	if err != nil {
		//response.UnauthorizedError(w, err.Error())
		return err // crate error 401
	}

	acGrant, err := c.useCase.AuthorizationCodeGrant(
		ctx.Request().Context(),
		aDto.Code,
		aDto.RedirectUri,
		aDto.ToModel(client.ID))

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, acGrant)
}

func (c controller) PasswordGrant(ctx echo.Context) error {
	aDto := new(oauthDto.GrantPasswordRequestDto)
	if err := ctx.Bind(aDto); err != nil {
		return oauthException.PasswordGrantBindingExc()
	}

	if err := aDto.ValidatePasswordDto(); err != nil {
		return oauthException.PasswordGrantValidationExc(err)
	}

	client, err := c.BasicAuthClient(ctx)
	if err != nil {
		//response.UnauthorizedError(w, err.Error())
		return err // crate error 401
	}

	acGrant, err := c.useCase.PasswordGrant(
		ctx.Request().Context(),
		aDto.Username,
		aDto.Password,
		aDto.Scope,
		aDto.ToModel(client.ID))

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, acGrant)
}

func (c controller) ClientCredentialsGrant(ctx echo.Context) error {
	aDto := new(oauthDto.GrantClientCredentialsRequestDto)
	if err := ctx.Bind(aDto); err != nil {
		return oauthException.GrantClientCredentialGrantBindingExc()
	}

	if err := aDto.ValidateClientCredentialsDto(); err != nil {
		return oauthException.AuthorizationCodeGrantValidationExc(err)
	}

	client, err := c.BasicAuthClient(ctx)
	if err != nil {
		//response.UnauthorizedError(w, err.Error())
		return err // crate error 401
	}

	acGrant, err := c.useCase.ClientCredentialsGrant(
		ctx.Request().Context(),
		aDto.Scope,
		aDto.ToModel(client.ID))

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, acGrant)
}

func (c controller) RefreshTokenGrant(ctx echo.Context) error {
	aDto := new(oauthDto.RefreshTokenRequestDto)
	if err := ctx.Bind(aDto); err != nil {
		return oauthException.RefreshTokenGrantBindingExc()
	}

	if err := aDto.ValidateRefreshTokenDto(); err != nil {
		return oauthException.RefreshTokenGrantValidationExc(err)
	}

	client, err := c.BasicAuthClient(ctx)
	if err != nil {
		//response.UnauthorizedError(w, err.Error())
		return err // crate error 401
	}

	acGrant, err := c.useCase.ClientCredentialsGrant(
		ctx.Request().Context(),
		aDto.Scope,
		aDto.ToModel(client.ID))

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, acGrant)
}

func (c controller) Introspect(ctx echo.Context) error {
	aDto := new(oauthDto.IntrospectRequestDto)
	if err := ctx.Bind(aDto); err != nil {
		return oauthException.IntrospectBindingExc()
	}

	if err := aDto.ValidateIntrospectDto(); err != nil {
		return oauthException.IntrospectValidationExc(err)
	}

	client, err := c.BasicAuthClient(ctx)
	if err != nil {
		//response.UnauthorizedError(w, err.Error())
		return err // crate error 401
	}

	article, err := c.useCase.IntrospectToken(
		ctx.Request().Context(),
		aDto.Token,
		aDto.TokenTypeHint,
		aDto.ToModel(client.ID))

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, article)
}

func (c controller) BasicAuthClient(ctx echo.Context) (*oauthModel.Client, error) {
	// Get client credentials from basic auth
	clientID, secret, ok := ctx.Request().BasicAuth()
	if !ok {
		return nil, errors.New("Invalid client ID or secret")
	}

	// Authenticate the client
	client, err := c.useCase.AuthClient(ctx.Request().Context(), clientID, secret)
	if err != nil {
		// For security reasons, return a general error message
		return nil, errors.New("Invalid client ID or secret")
	}

	return client, nil
}
