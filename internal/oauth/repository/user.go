package oauthRepository

import (
	"database/sql"
	"errors"
	"fmt"
	oauthDomain "github.com/diki-haryadi/go-micro-template/internal/oauth/domain/model"
	"github.com/diki-haryadi/go-micro-template/pkg"
	"github.com/google/uuid"
	"strings"
	"time"
)

var (
	// MinPasswordLength defines minimum password length
	MinPasswordLength = 6

	// ErrPasswordTooShort ...
	ErrPasswordTooShort = fmt.Errorf(
		"Password must be at least %d characters long",
		MinPasswordLength,
	)
	// ErrUserNotFound ...
	ErrUserNotFound = errors.New("User not found")
	// ErrInvalidUserPassword ...
	ErrInvalidUserPassword = errors.New("Invalid user password")
	// ErrCannotSetEmptyUsername ...
	ErrCannotSetEmptyUsername = errors.New("Cannot set empty username")
	// ErrUserPasswordNotSet ...
	ErrUserPasswordNotSet = errors.New("User password not set")
	// ErrUsernameTaken ...
	ErrUsernameTaken = errors.New("Username taken")
)

// FindUserByUsername looks up a user by username
// FindUserByUsername retrieves a user by their username using raw SQL
func (rp *repository) FindUserByUsername(username string) (*oauthDomain.Users, error) {
	// Prepare the SQL query to fetch the user by case-insensitive username
	sqlQuery := "SELECT id, username, password, role_id, created_at, updated_at FROM users WHERE LOWER(username) = $1"

	// Execute the query
	user := new(oauthDomain.Users)
	err := rp.postgres.SqlxDB.QueryRow(sqlQuery, strings.ToLower(username)).Scan(
		&user.ID, &user.Username, &user.Password, &user.RoleID, &user.CreatedAt, &user.UpdatedAt,
	)

	// Check if no rows are found (user not found)
	if err == sql.ErrNoRows {
		return nil, ErrUserNotFound
	}
	if err != nil {
		return nil, err // Handle any other error
	}

	return user, nil
}

// CreateUserCommon creates a new user using raw SQL
func (rp *repository) CreateUserCommon(roleID, username, password string) (*oauthDomain.Users, error) {
	// Start with a user struct, setting initial fields
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
		if len(password) < MinPasswordLength {
			return nil, ErrPasswordTooShort
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
	err := rp.postgres.SqlxDB.QueryRow(sqlCheckUsername, user.Username).Scan(&count)
	if err != nil {
		return nil, err
	}

	if count > 0 {
		return nil, ErrUsernameTaken
	}

	// Insert the new user into the database
	sqlInsert := `
        INSERT INTO users (id, created_at, role_id, username, password)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id, created_at, role_id, username, password
    `
	err = rp.postgres.SqlxDB.QueryRow(sqlInsert, user.ID, user.CreatedAt, user.RoleID, user.Username, user.Password).
		Scan(&user.ID, &user.CreatedAt, &user.RoleID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// SetPasswordCommon updates the user's password using raw SQL
func (rp *repository) SetPasswordCommon(user *oauthDomain.Users, password string) error {
	if len(password) < MinPasswordLength {
		return ErrPasswordTooShort
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
	_, err = rp.postgres.SqlxDB.Exec(sqlQuery, string(passwordHash), time.Now().UTC(), user.ID)
	if err != nil {
		return err
	}

	return nil
}

// UpdateUsernameCommon updates the user's username using raw SQL
func (rp *repository) UpdateUsernameCommon(user *oauthDomain.Users, username string) error {
	if username == "" {
		return ErrCannotSetEmptyUsername
	}

	// Prepare the SQL query to update the username field
	sqlQuery := `
        UPDATE users
        SET username = $1
        WHERE id = $2`

	// Execute the query to update the username
	_, err := rp.postgres.SqlxDB.Exec(sqlQuery, strings.ToLower(username), user.ID)
	if err != nil {
		return err
	}

	return nil
}
