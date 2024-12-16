package oauthUseCase

import (
	"context"
	oauthDto "github.com/diki-haryadi/go-micro-template/internal/oauth/dto"
	"github.com/diki-haryadi/go-micro-template/pkg"
	"github.com/diki-haryadi/go-micro-template/pkg/response"
)

func (uc *useCase) ForgotPassword(ctx context.Context, dto *oauthDto.ForgotPasswordRequest) (*oauthDto.ForgotPasswordResponse, error) {
	if len(dto.Password) < response.MinPasswordLength {
		return &oauthDto.ForgotPasswordResponse{}, response.ErrPasswordTooShort
	}

	user, err := uc.repository.FetchUserByUserID(ctx, dto.UUID)
	if err != nil {
		return &oauthDto.ForgotPasswordResponse{}, nil
	}

	err = pkg.VerifyPassword(user.Password.String, dto.Password)
	if err == nil {
		return &oauthDto.ForgotPasswordResponse{}, response.ErrInvalidPasswordCannotSame
	}

	passwordHash, err := pkg.HashPassword(dto.Password)
	if err != nil {
		return &oauthDto.ForgotPasswordResponse{}, err
	}

	err = uc.repository.UpdatePassword(ctx, dto.UUID, string(passwordHash))
	if err != nil {
		return &oauthDto.ForgotPasswordResponse{}, nil
	}

	return &oauthDto.ForgotPasswordResponse{
		Status: true,
	}, nil
}
