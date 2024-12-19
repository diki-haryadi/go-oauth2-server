package authHttpController

import (
	oauthDomain "github.com/diki-haryadi/go-micro-template/internal/authentication/domain"
	oauthModel "github.com/diki-haryadi/go-micro-template/internal/authentication/domain/model"
	oauthDto "github.com/diki-haryadi/go-micro-template/internal/authentication/dto"
	"github.com/diki-haryadi/go-micro-template/pkg/response"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type controller struct {
	useCase oauthDomain.UseCase
}

func NewController(uc oauthDomain.UseCase) oauthDomain.HttpController {
	return &controller{
		useCase: uc,
	}
}

func (c controller) BasicAuthClient(ctx echo.Context) (*oauthModel.Client, error) {
	// Get client credentials from basic auth
	clientID, secret, ok := ctx.Request().BasicAuth()
	if !ok {
		return nil, response.ErrInvalidClientIDOrSecret
	}

	// Authenticate the client
	client, err := c.useCase.AuthClient(ctx.Request().Context(), clientID, secret)
	if err != nil {
		return nil, response.ErrInvalidClientIDOrSecret
	}

	return client, nil
}

func (c controller) Register(ctx echo.Context) error {
	res := response.NewJSONResponse()
	aDto := new(oauthDto.UserRequestDto).GetFieldsUser(ctx)
	if err := aDto.ValidateUserDto(); err != nil {
		res.SetError(response.ErrInvalidIntrospectRequest).SetMessage(err.Error()).Send(ctx.Response().Writer)
		return nil
	}

	_, err := c.BasicAuthClient(ctx)
	if err != nil {
		res.SetError(err).SetMessage(err.Error()).Send(ctx.Response().Writer)
		return nil
	}

	register, err := c.useCase.Register(
		ctx.Request().Context(),
		aDto)

	if err != nil {
		res.SetError(err).SetMessage(err.Error()).Send(ctx.Response().Writer)
		return nil
	}

	res.APIStatusSuccess().SetData(*register).Send(ctx.Response().Writer)
	return nil
}

func (c controller) ChangePassword(ctx echo.Context) error {
	res := response.NewJSONResponse()
	aDto := new(oauthDto.ChangePasswordRequest).GetFieldsChangePassword(ctx)
	if err := aDto.ValidateChangePasswordDto(); err != nil {
		res.SetError(response.ErrInvalidIntrospectRequest).SetMessage(err.Error()).Send(ctx.Response().Writer)
		return nil
	}

	_, err := c.BasicAuthClient(ctx)
	if err != nil {
		res.SetError(err).SetMessage(err.Error()).Send(ctx.Response().Writer)
		return nil
	}

	changePass, err := c.useCase.ChangePassword(
		ctx.Request().Context(), aDto)

	if err != nil {
		res.SetError(err).SetMessage(err.Error()).Send(ctx.Response().Writer)
		return nil
	}

	res.APIStatusSuccess().SetData(*changePass).Send(ctx.Response().Writer)
	return nil
}

func (c controller) ForgotPassword(ctx echo.Context) error {
	res := response.NewJSONResponse()
	aDto := new(oauthDto.ForgotPasswordRequest).GetFieldsForgotPassword(ctx)
	if err := aDto.ValidateForgotPasswordDto(); err != nil {
		res.SetError(response.ErrInvalidIntrospectRequest).SetMessage(err.Error()).Send(ctx.Response().Writer)
		return nil
	}

	_, err := c.BasicAuthClient(ctx)
	if err != nil {
		res.SetError(err).SetMessage(err.Error()).Send(ctx.Response().Writer)
		return nil
	}

	forgotPass, err := c.useCase.ForgotPassword(
		ctx.Request().Context(), aDto)

	if err != nil {
		res.SetError(err).SetMessage(err.Error()).Send(ctx.Response().Writer)
		return nil
	}

	res.APIStatusSuccess().SetData(*forgotPass).Send(ctx.Response().Writer)
	return nil
}

func (c controller) UpdateUsername(ctx echo.Context) error {
	res := response.NewJSONResponse()
	aDto := new(oauthDto.UpdateUsernameRequest).GetFieldsUpdateUsername(ctx)
	if err := aDto.ValidateUsernameDto(); err != nil {
		res.SetError(response.ErrInvalidIntrospectRequest).SetMessage(err.Error()).Send(ctx.Response().Writer)
		return nil
	}

	_, err := c.BasicAuthClient(ctx)
	if err != nil {
		res.SetError(err).SetMessage(err.Error()).Send(ctx.Response().Writer)
		return nil
	}

	uuid, err := uuid.Parse(aDto.UUID)
	if err != nil {
		res.SetError(err).SetMessage(err.Error()).Send(ctx.Response().Writer)
		return nil
	}

	err = c.useCase.UpdateUsername(
		ctx.Request().Context(),
		aDto.ToModel(uuid), aDto.Username)

	if err != nil {
		res.SetError(err).SetMessage(err.Error()).Send(ctx.Response().Writer)
		return nil
	}

	res.APIStatusSuccess().Send(ctx.Response().Writer)
	return nil
}
