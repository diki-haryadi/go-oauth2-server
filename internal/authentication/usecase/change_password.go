package authUseCase

import (
	"context"
	authDto "github.com/diki-haryadi/go-micro-template/internal/authentication/dto"
	"github.com/diki-haryadi/go-micro-template/pkg"
	"github.com/diki-haryadi/go-micro-template/pkg/response"
)

func (uc *useCase) ChangePassword(ctx context.Context, dto *authDto.ChangePasswordRequest) (*authDto.ChangePasswordResponse, error) {
	if len(dto.Password) < response.MinPasswordLength {
		return &authDto.ChangePasswordResponse{}, response.ErrPasswordTooShort
	}

	user, err := uc.repository.FetchUserByUserID(ctx, dto.UUID)
	if err != nil {
		return &authDto.ChangePasswordResponse{}, nil
	}

	err = pkg.VerifyPassword(user.Password.String, dto.Password)
	if err != nil {
		return &authDto.ChangePasswordResponse{}, response.ErrInvalidPassword
	}

	passwordHash, err := pkg.HashPassword(dto.NewPassword)
	if err != nil {
		return &authDto.ChangePasswordResponse{}, err
	}

	err = uc.repository.UpdatePassword(ctx, dto.UUID, string(passwordHash))
	if err != nil {
		return &authDto.ChangePasswordResponse{}, nil
	}

	return &authDto.ChangePasswordResponse{
		Status: true,
	}, nil
}
