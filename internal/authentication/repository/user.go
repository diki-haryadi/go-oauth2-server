package authRepository

import (
	"context"
	"database/sql"
	oauthDomain "github.com/diki-haryadi/go-micro-template/internal/authentication/domain/model"
	"github.com/diki-haryadi/go-micro-template/pkg"
	"github.com/diki-haryadi/go-micro-template/pkg/response"
	"github.com/google/uuid"
	"strings"
	"time"
)

func (rp *repository) FindUserByUsername(ctx context.Context, username string) (*oauthDomain.Users, error) {
	sqlQuery := "SELECT id, username, password, role_id, created_at, updated_at FROM users WHERE LOWER(username) = $1"

	user := new(oauthDomain.Users)
	err := rp.postgres.SqlxDB.QueryRowContext(ctx, sqlQuery, strings.ToLower(username)).Scan(
		&user.ID, &user.Username, &user.Password, &user.RoleID, &user.CreatedAt, &user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, response.ErrUserNotFound
	}
	if err != nil {
		return nil, err // Handle any other error
	}

	return user, nil
}

func (rp *repository) CreateUserCommon(ctx context.Context, roleID, username, password string) (*oauthDomain.Users, error) {
	user := &oauthDomain.Users{
		Common: oauthDomain.Common{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
		},
		RoleID:   pkg.StringOrNull(roleID),
		Username: strings.ToLower(username),
		Password: pkg.StringOrNull(""),
	}

	// If the password is being set, hash it
	if password != "" {
		if len(password) < response.MinPasswordLength {
			return nil, response.ErrPasswordTooShort
		}

		passwordHash, err := pkg.HashPassword(password)
		if err != nil {
			return nil, err
		}
		user.Password = pkg.StringOrNull(string(passwordHash))
	}

	// Check if the username is already taken using raw SQL
	sqlCheckUsername := "SELECT COUNT(*) FROM users WHERE LOWER(username) = $1"
	var count int
	err := rp.postgres.SqlxDB.QueryRowContext(ctx, sqlCheckUsername, user.Username).Scan(&count)
	if err != nil {
		return nil, err
	}

	if count > 0 {
		return nil, response.ErrUsernameTaken
	}

	// Insert the new user into the database
	sqlInsert := `
        INSERT INTO users (id, created_at, role_id, username, password)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id, created_at, role_id, username, password
    `
	err = rp.postgres.SqlxDB.QueryRowContext(ctx, sqlInsert, user.ID, user.CreatedAt, user.RoleID, user.Username, user.Password).
		Scan(&user.ID, &user.CreatedAt, &user.RoleID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// SetPasswordCommon updates the user's password using raw SQL
func (rp *repository) SetPasswordCommon(ctx context.Context, user *oauthDomain.Users, password string) error {
	if len(password) < response.MinPasswordLength {
		return response.ErrPasswordTooShort
	}

	// Create a bcrypt hash for the password
	passwordHash, err := pkg.HashPassword(password)
	if err != nil {
		return err
	}

	// Prepare the SQL query to update the password and the updated_at field
	sqlQuery := `
        UPDATE users
        SET password = $1, updated_at = $2
        WHERE id = $3
    `

	// Execute the query to update the user's password
	_, err = rp.postgres.SqlxDB.ExecContext(ctx, sqlQuery, string(passwordHash), time.Now().UTC(), user.ID)
	if err != nil {
		return err
	}

	return nil
}

// UpdateUsernameCommon updates the user's username using raw SQL
func (rp *repository) UpdateUsernameCommon(ctx context.Context, user *oauthDomain.Users, username string) error {
	if username == "" {
		return response.ErrCannotSetEmptyUsername
	}

	// Prepare the SQL query to update the username field
	sqlQuery := `
        UPDATE users
        SET username = $1
        WHERE id = $2`

	// Execute the query to update the username
	_, err := rp.postgres.SqlxDB.ExecContext(ctx, sqlQuery, strings.ToLower(username), user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (rp *repository) UpdatePassword(ctx context.Context, uuid, password string) error {
	if password == "" {
		return response.ErrUserPasswordNotSet
	}

	// Prepare the SQL query to update the username field
	sqlQuery := `
        UPDATE users
        SET password = $1
        WHERE id = $2`

	// Execute the query to update the username
	_, err := rp.postgres.SqlxDB.Exec(sqlQuery, password, uuid)
	if err != nil {
		return err
	}

	return nil
}

// FetchUserByUserID retrieves the user by user_id using raw SQL
func (rp *repository) FetchUserByUserID(ctx context.Context, userID string) (*oauthDomain.Users, error) {
	sqlUserQuery := "SELECT id, username, password FROM users WHERE id = $1"
	user := new(oauthDomain.Users)
	err := rp.postgres.SqlxDB.QueryRowContext(ctx, sqlUserQuery, userID).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, response.ErrUserNotFound
		}
		return nil, err
	}
	return user, nil
}
