package oauthUseCase

import (
	"context"
	"errors"
)

var (
	// ErrInvalidScope ...
	ErrInvalidScope = errors.New("Invalid scope")
)

func (uc *useCase) GetScope(ctx context.Context, requestScope string) (string, error) {
	if requestScope == "" {
		scope := uc.repository.GetDefaultScope(ctx)
		return scope, nil
	}

	if scope := uc.repository.ScopeExists(ctx, requestScope); scope {
		return requestScope, nil
	}
	return "", ErrInvalidScope
}
