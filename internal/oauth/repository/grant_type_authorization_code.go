package oauthRepository

import (
	"context"
	"database/sql"
	oauthDomain "github.com/diki-haryadi/go-micro-template/internal/oauth/domain/model"
	"github.com/diki-haryadi/go-micro-template/pkg/response"
)

// FetchAuthorizationCodeByCode retrieves the authorization code from the database using raw SQL
func (rp *repository) FetchAuthorizationCodeByCode(ctx context.Context, client *oauthDomain.Client, code string) (*oauthDomain.AuthorizationCode, error) {
	sqlQuery := `
        SELECT id, client_id, user_id, code, redirect_uri, expires_at, scope
        FROM authorization_codes
        WHERE client_id = $1 AND code = $2
    `
	var authorizationCode oauthDomain.AuthorizationCode
	row := rp.postgres.SqlxDB.QueryRow(sqlQuery, client.ID, code)

	// Scan the result into the authorizationCode struct
	err := row.Scan(
		&authorizationCode.ID,
		&authorizationCode.ClientID,
		&authorizationCode.UserID,
		&authorizationCode.Code,
		&authorizationCode.RedirectURI,
		&authorizationCode.ExpiresAt,
		&authorizationCode.Scope,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, response.ErrAuthorizationCodeNotFound
		}
		return nil, err
	}

	return &authorizationCode, nil
}

// DeleteAuthorizationCode deletes the authorization code from the database after use
func (rp *repository) DeleteAuthorizationCode(authorizationCodeID string) error {
	sqlDelete := `
        DELETE FROM authorization_codes WHERE id = $1
    `
	_, err := rp.postgres.SqlxDB.Exec(sqlDelete, authorizationCodeID)
	return err
}
