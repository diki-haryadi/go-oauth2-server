package oauthUseCase

import (
	"context"
	oauthDto "github.com/diki-haryadi/go-micro-template/internal/oauth/dto"
	"github.com/diki-haryadi/go-micro-template/pkg"
	"github.com/diki-haryadi/go-micro-template/pkg/response"
)

func (uc *useCase) ChangePassword(ctx context.Context, dto *oauthDto.ChangePasswordRequest) (*oauthDto.ChangePasswordResponse, error) {
	if len(dto.Password) < response.MinPasswordLength {
		return &oauthDto.ChangePasswordResponse{}, response.ErrPasswordTooShort
	}

	user, err := uc.repository.FetchUserByUserID(ctx, dto.UUID)
	if err != nil {
		return &oauthDto.ChangePasswordResponse{}, nil
	}

	err = pkg.VerifyPassword(user.Password.String, dto.Password)
	if err != nil {
		return &oauthDto.ChangePasswordResponse{}, response.ErrInvalidPassword
	}

	passwordHash, err := pkg.HashPassword(dto.NewPassword)
	if err != nil {
		return &oauthDto.ChangePasswordResponse{}, err
	}

	err = uc.repository.UpdatePassword(ctx, dto.UUID, string(passwordHash))
	if err != nil {
		return &oauthDto.ChangePasswordResponse{}, nil
	}

	return &oauthDto.ChangePasswordResponse{
		Status: true,
	}, nil
}
