package oauthUseCase

import (
	"context"
	"github.com/diki-haryadi/go-micro-template/internal/oauth/dto"
)

func (uc *useCase) Register(ctx context.Context, dto *oauthDto.UserRequestDto) (*oauthDto.UserResponse, error) {
	user, err := uc.repository.CreateUserCommon(ctx, dto.RoleID, dto.Username, dto.Password)
	if err != nil {
		return &oauthDto.UserResponse{}, nil
	}
	return &oauthDto.UserResponse{
		Username: user.Username,
		Role:     user.Role.Name,
	}, nil
}
