package oauthUseCase

import (
	"context"
	oauthDomain "github.com/diki-haryadi/go-micro-template/internal/oauth/domain/model"
)

func (uc *useCase) GrantAccessToken(ctx context.Context, client *oauthDomain.Client, user *oauthDomain.Users, expiresIn int, scope string) (*oauthDomain.AccessToken, error) {
	accessToken, err := uc.repository.GrantAccessToken(ctx, client, user, expiresIn, scope)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}
