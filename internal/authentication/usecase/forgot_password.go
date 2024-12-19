package authUseCase

import (
	"context"
	authDto "github.com/diki-haryadi/go-micro-template/internal/authentication/dto"
	"github.com/diki-haryadi/go-micro-template/pkg"
	"github.com/diki-haryadi/go-micro-template/pkg/response"
)

func (uc *useCase) ForgotPassword(ctx context.Context, dto *authDto.ForgotPasswordRequest) (*authDto.ForgotPasswordResponse, error) {
	if len(dto.Password) < response.MinPasswordLength {
		return &authDto.ForgotPasswordResponse{}, response.ErrPasswordTooShort
	}

	user, err := uc.repository.FetchUserByUserID(ctx, dto.UUID)
	if err != nil {
		return &authDto.ForgotPasswordResponse{}, nil
	}

	err = pkg.VerifyPassword(user.Password.String, dto.Password)
	if err == nil {
		return &authDto.ForgotPasswordResponse{}, response.ErrInvalidPasswordCannotSame
	}

	passwordHash, err := pkg.HashPassword(dto.Password)
	if err != nil {
		return &authDto.ForgotPasswordResponse{}, err
	}

	err = uc.repository.UpdatePassword(ctx, dto.UUID, string(passwordHash))
	if err != nil {
		return &authDto.ForgotPasswordResponse{}, nil
	}

	return &authDto.ForgotPasswordResponse{
		Status: true,
	}, nil
}
