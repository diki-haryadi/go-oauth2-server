package oauthUseCase

import (
	"context"
	oauthDomain "github.com/diki-haryadi/go-micro-template/internal/oauth/domain/model"
	"github.com/diki-haryadi/go-micro-template/pkg"
	"github.com/diki-haryadi/go-micro-template/pkg/response"
)

// UserExists returns true if user exists
func (uc *useCase) UserExists(ctx context.Context, username string) bool {
	_, err := uc.repository.FindUserByUsername(ctx, username)
	return err == nil
}

// CreateUser saves a new user to database
func (uc *useCase) CreateUser(ctx context.Context, roleID, username, password string) (*oauthDomain.Users, error) {
	return uc.repository.CreateUserCommon(ctx, roleID, username, password)
}

// CreateUserTx saves a new user to database using injected db object
func (uc *useCase) CreateUserTx(ctx context.Context, roleID, username, password string) (*oauthDomain.Users, error) {
	return uc.repository.CreateUserCommon(ctx, roleID, username, password)
}

// SetPassword sets a user password
func (uc *useCase) SetPassword(ctx context.Context, user *oauthDomain.Users, password string) error {
	return uc.repository.SetPasswordCommon(ctx, user, password)
}

// SetPasswordTx sets a user password in a transaction
func (uc *useCase) SetPasswordTx(ctx context.Context, user *oauthDomain.Users, password string) error {
	return uc.repository.SetPasswordCommon(ctx, user, password)
}

// AuthUser authenticates user
func (uc *useCase) AuthUser(ctx context.Context, username, password string) (*oauthDomain.Users, error) {
	// Fetch the user
	user, err := uc.repository.FindUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	// Check that the password is set
	if !user.Password.Valid {
		return nil, err
	}

	// Verify the password
	if pkg.VerifyPassword(user.Password.String, password) != nil {
		return nil, err
	}

	role, err := uc.repository.FindRoleByID(ctx, user.RoleID.String)
	if err != nil {
		return nil, err
	}
	user.Role = role
	return user, nil
}

// UpdateUsername ...
func (uc *useCase) UpdateUsername(ctx context.Context, user *oauthDomain.Users, username string) error {
	if username == "" {
		return response.ErrCannotSetEmptyUsername
	}

	return uc.repository.UpdateUsernameCommon(ctx, user, username)
}

// UpdateUsernameTx ...
func (uc *useCase) UpdateUsernameTx(ctx context.Context, user *oauthDomain.Users, username string) error {
	return uc.repository.UpdateUsernameCommon(ctx, user, username)
}
