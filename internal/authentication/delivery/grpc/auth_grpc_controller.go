package authGrpcController

import (
	"context"
	oauthDomain "github.com/diki-haryadi/go-micro-template/internal/authentication/domain"
	oauthModel "github.com/diki-haryadi/go-micro-template/internal/authentication/domain/model"
	oauthDto "github.com/diki-haryadi/go-micro-template/internal/authentication/dto"
	"github.com/diki-haryadi/go-micro-template/pkg/response"
	auhenticationV1 "github.com/diki-haryadi/protobuf-ecomerce/oauth2_server_service/authentication/v1"
	"github.com/google/uuid"
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

func (c *controller) Register(ctx context.Context, req *auhenticationV1.RegisterRequest) (*auhenticationV1.RegisterResponse, error) {
	aDto := new(oauthDto.UserRequestDto).GetFieldsUserValue(req.Username, req.Password, req.RoleId)
	if err := aDto.ValidateUserDto(); err != nil {
		return &auhenticationV1.RegisterResponse{}, err
	}

	_, err := c.BasicAuthClient(ctx, req.ClientId, req.ClientSecret)
	if err != nil {
		return &auhenticationV1.RegisterResponse{}, err
	}

	register, err := c.useCase.Register(
		ctx,
		aDto)

	if err != nil {
		return &auhenticationV1.RegisterResponse{}, err
	}

	return &auhenticationV1.RegisterResponse{
		Username: register.Username,
		Role:     register.Role,
	}, err
}

func (c *controller) ChangePassword(ctx context.Context, req *auhenticationV1.ChangePasswordRequest) (*auhenticationV1.ChangePasswordResponse, error) {
	aDto := new(oauthDto.ChangePasswordRequest).GetFieldsChangePasswordValue(req.Uuid, req.Password, req.NewPassword)
	if err := aDto.ValidateChangePasswordDto(); err != nil {
		return &auhenticationV1.ChangePasswordResponse{}, err
	}

	_, err := c.BasicAuthClient(ctx, req.ClientId, req.ClientSecret)
	if err != nil {
		return &auhenticationV1.ChangePasswordResponse{}, err
	}

	forgotPass, err := c.useCase.ChangePassword(
		ctx, aDto)

	if err != nil {
		return &auhenticationV1.ChangePasswordResponse{}, err
	}
	return &auhenticationV1.ChangePasswordResponse{
		Status: forgotPass.Status,
	}, err
}

func (c *controller) ForgotPassword(ctx context.Context, req *auhenticationV1.ForgotPasswordRequest) (*auhenticationV1.ForgotPasswordResponse, error) {
	aDto := new(oauthDto.ForgotPasswordRequest).GetFieldsForgotPasswordValue(req.Uuid, req.Password)
	if err := aDto.ValidateForgotPasswordDto(); err != nil {
		return &auhenticationV1.ForgotPasswordResponse{}, err
	}

	_, err := c.BasicAuthClient(ctx, req.ClientId, req.ClientSecret)
	if err != nil {
		return &auhenticationV1.ForgotPasswordResponse{}, err
	}

	forgotPass, err := c.useCase.ForgotPassword(
		ctx, aDto)

	if err != nil {
		return &auhenticationV1.ForgotPasswordResponse{}, err
	}
	return &auhenticationV1.ForgotPasswordResponse{
		Status: forgotPass.Status,
	}, err
}

func (c *controller) UpdateUsername(ctx context.Context, req *auhenticationV1.UpdateUsernameRequest) (*auhenticationV1.UpdateUsernameResponse, error) {
	aDto := new(oauthDto.UpdateUsernameRequest).GetFieldsUpdateUsernameValue(req.Uuid, req.Username)
	if err := aDto.ValidateUsernameDto(); err != nil {
		return &auhenticationV1.UpdateUsernameResponse{}, err
	}

	_, err := c.BasicAuthClient(ctx, req.ClientId, req.ClientSecret)
	if err != nil {
		return &auhenticationV1.UpdateUsernameResponse{}, err
	}

	uuid, err := uuid.Parse(aDto.UUID)
	if err != nil {
		return &auhenticationV1.UpdateUsernameResponse{}, err
	}

	err = c.useCase.UpdateUsername(
		ctx,
		aDto.ToModel(uuid), aDto.Username)

	if err != nil {
		return &auhenticationV1.UpdateUsernameResponse{}, err
	}

	return &auhenticationV1.UpdateUsernameResponse{}, err
}
