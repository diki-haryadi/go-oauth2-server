package oauthUseCase

import (
	"context"
	"errors"
	"github.com/diki-haryadi/go-micro-template/internal/oauth/domain/model"
	"github.com/diki-haryadi/go-micro-template/pkg"
)

var (
	// ErrClientNotFound ...
	ErrClientNotFound = errors.New("Client not found")
	// ErrInvalidClientSecret ...
	ErrInvalidClientSecret = errors.New("Invalid client secret")
)

func (uc *useCase) AuthClient(ctx context.Context, clientID, secret string) (*oauthDomain.Client, error) {
	client, err := uc.repository.FindClientByClientID(ctx, clientID)
	if err != nil {
		return nil, ErrClientNotFound
	}

	if pkg.VerifyPassword(client.Secret, secret) != nil {
		return nil, ErrInvalidClientSecret
	}
	return client, nil
}

func (uc *useCase) CreateClient(ctx context.Context, clientID, secret, redirectURI string) (*oauthDomain.Client, error) {
	client, err := uc.repository.CreateClientCommon(ctx, clientID, secret, redirectURI)
	if err != nil {
		return nil, err
	}
	return client, err
}

func (uc *useCase) ClientExists(ctx context.Context, clientID string) bool {
	_, err := uc.repository.FindClientByClientID(ctx, clientID)
	return err == nil
}
