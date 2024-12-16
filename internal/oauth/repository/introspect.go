package oauthRepository

import (
	"context"
	"database/sql"
	oauthDomain "github.com/diki-haryadi/go-micro-template/internal/oauth/domain/model"
	"github.com/diki-haryadi/go-micro-template/pkg/response"
)

// FetchClientByClientID retrieves the client by client_id using raw SQL
func (rp *repository) FetchClientByClientID(ctx context.Context, clientID string) (*oauthDomain.Client, error) {
	sqlClientQuery := "SELECT key FROM clients WHERE id = $1"
	client := new(oauthDomain.Client)
	err := rp.postgres.SqlxDB.QueryRowContext(ctx, sqlClientQuery, clientID).Scan(&client.Key)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, response.ErrClientNotFound
		}
		return nil, err
	}
	return client, nil
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
