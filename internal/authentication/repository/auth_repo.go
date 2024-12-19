package authRepository

import (
	oauthDomain "github.com/diki-haryadi/go-micro-template/internal/authentication/domain"
	"github.com/diki-haryadi/ztools/postgres"
)

type repository struct {
	postgres *postgres.Postgres
}

func NewRepository(conn *postgres.Postgres) oauthDomain.Repository {
	return &repository{postgres: conn}
}
