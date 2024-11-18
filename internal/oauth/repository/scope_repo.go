package oauthRepository

import (
	"context"
	"errors"
	"sort"
	"strconv"
	"strings"
)

var (
	// ErrInvalidScope ...
	ErrInvalidScope = errors.New("Invalid scope")
)

// GetScope takes a requested scope and, if it's empty, returns the default
// scope, if not empty, it validates the requested scope
func (rp *repository) GetScope(ctx context.Context, requestedScope string) (string, error) {
	// Return the default scope if the requested scope is empty
	if requestedScope == "" {
		return rp.GetDefaultScope(ctx), nil
	}

	// If the requested scope exists in the database, return it
	if rp.ScopeExists(ctx, requestedScope) {
		return requestedScope, nil
	}

	// Otherwise return error
	return "", ErrInvalidScope
}

// GetDefaultScope retrieves the default scope from the database using raw SQL
func (rp *repository) GetDefaultScope(ctx context.Context) string {
	// Fetch default scopes from the database using raw SQL
	sqlQuery := "SELECT scope FROM scopes WHERE is_default = $1"
	rows, err := rp.postgres.SqlxDB.Query(sqlQuery, true)
	if err != nil {
		// Handle error (e.g., database connection issues)
		return ""
	}
	defer rows.Close()

	var scopes []string
	for rows.Next() {
		var scope string
		if err := rows.Scan(&scope); err != nil {
			// Handle error (e.g., scanning issues)
			return ""
		}
		scopes = append(scopes, scope)
	}

	// Sort the scopes alphabetically
	sort.Strings(scopes)

	// Return space-delimited scope string
	return strings.Join(scopes, " ")
}

// ScopeExists checks if a scope exists using raw SQL
func (rp *repository) ScopeExists(ctx context.Context, requestedScope string) bool {
	// Split the requested scope string
	scopes := strings.Split(requestedScope, " ")

	// Prepare a query to count how many of the requested scopes exist in the database
	query := "SELECT COUNT(*) FROM scopes WHERE scope IN ("
	// Build the query with placeholders for each scope
	placeholders := make([]string, len(scopes))
	for i := range scopes {
		placeholders[i] = "$" + strconv.Itoa(i+1) // Generate placeholders for parameterized queries
	}
	query += strings.Join(placeholders, ", ") + ")"

	// Execute the query and get the count
	var count int
	err := rp.postgres.SqlxDB.QueryRow(query, scopes).Scan(&count)
	if err != nil {
		// Handle error (e.g., database connection issues)
		return false
	}

	// Return true only if all requested scopes are found
	return count == len(scopes)
}
