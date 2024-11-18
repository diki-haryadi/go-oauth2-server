package oauthRepository

import (
	"database/sql"
	"errors"
	oauthDomain "github.com/diki-haryadi/go-micro-template/internal/oauth/domain/model"
)

var (
	// ErrRoleNotFound ...
	ErrRoleNotFound = errors.New("Role not found")
)

// FindRoleByID looks up a role by ID and returns it
// FindRoleByID retrieves a role by its ID using raw SQL
func (rp *repository) FindRoleByID(id string) (*oauthDomain.Role, error) {
	// Define the SQL query to fetch the role by ID
	sqlQuery := "SELECT id, name, description FROM roles WHERE id = $1"

	// Execute the query and scan the result into the role object
	role := new(oauthDomain.Role)
	err := rp.postgres.SqlxDB.QueryRow(sqlQuery, id).Scan(&role.ID, &role.Name, &role.Description)

	// Handle case where no rows are returned (role not found)
	if err == sql.ErrNoRows {
		return nil, ErrRoleNotFound
	}
	if err != nil {
		return nil, err // Other errors, such as database connection issues
	}

	return role, nil
}
