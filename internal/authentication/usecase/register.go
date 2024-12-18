package authUseCase

import (
	"context"
	"github.com/diki-haryadi/go-micro-template/internal/authentication/dto"
)

func (uc *useCase) Register(ctx context.Context, dto *authDto.UserRequestDto) (*authDto.UserResponse, error) {
	user, err := uc.repository.CreateUserCommon(ctx, dto.RoleID, dto.Username, dto.Password)
	if err != nil {
		return &authDto.UserResponse{}, nil
	}
	return &authDto.UserResponse{
		UUID:     user.ID.String(),
		Username: user.Username,
		RoleID:   user.RoleID.String,
	}, nil
}
