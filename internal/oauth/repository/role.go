package oauthRepository

import (
	"context"
	"database/sql"
	oauthDomain "github.com/diki-haryadi/go-micro-template/internal/oauth/domain/model"
	"github.com/diki-haryadi/go-micro-template/pkg/response"
)

// FindRoleByID retrieves a role by its ID using raw SQL
func (rp *repository) FindRoleByID(ctx context.Context, id string) (*oauthDomain.Role, error) {
	sqlQuery := "SELECT id, name FROM roles WHERE id = $1"

	role := new(oauthDomain.Role)
	err := rp.postgres.SqlxDB.QueryRowContext(ctx, sqlQuery, id).Scan(&role.ID, &role.Name)

	if err == sql.ErrNoRows {
		return nil, response.ErrRoleNotFound
	}
	if err != nil {
		return nil, err
	}

	return role, nil
}
