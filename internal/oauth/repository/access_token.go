package oauthRepository

import (
	"context"
	"fmt"
	oauthDomain "github.com/diki-haryadi/go-micro-template/internal/oauth/domain/model"
	oauthDto "github.com/diki-haryadi/go-micro-template/internal/oauth/dto"
)

func (rp *repository) GrantAccessToken(client *oauthDomain.Client, user *oauthDomain.Users, expiresIn int, scope string) (*oauthDomain.AccessToken, error) {
	// Begin a transaction
	tx, _ := rp.postgres.SqlxDB.BeginTx(context.Background(), nil)

	// Delete expired access tokens
	query, _ := tx.Query("SELECT * FROM access_tokens WHERE client_id = $1", client.ID)
	if user != nil && fmt.Sprint(user.ID) != "" {
		query, _ = tx.Query("SELECT * FROM access_tokens WHERE client_id = $1", user.ID)
	} else {
		query, _ = tx.Query("SELECT * FROM access_tokens WHERE user_id IS NULL")
	}
	fmt.Println(query)

	sqlQuery := "DELETE FROM access_tokens WHERE expires_at <= NOW()"
	// Execute the query
	_, err := tx.Exec(sqlQuery)
	if err != nil {
		// If an error occurs, rollback the transaction
		_ = tx.Rollback()
		return nil, err
	}

	// Create a new access token
	accessToken := oauthDto.NewOauthAccessToken(client, user, expiresIn, scope)
	sqlQueryAT := `
    INSERT INTO access_tokens (client, user_id, expires_in, scope)
    VALUES ($1, $2, $3, $4)`

	// Execute the INSERT query with the provided values
	_, err = tx.Exec(sqlQueryAT, client, user, expiresIn, scope)
	if err != nil {
		// If an error occurs, rollback the transaction
		_ = tx.Rollback()
		return nil, err
	}
	accessToken.Client = client
	accessToken.User = user

	// Commit the transaction
	if err = tx.Commit(); err != nil {
		_ = tx.Rollback() // rollback the transaction
		return nil, err
	}

	return accessToken, nil
}
