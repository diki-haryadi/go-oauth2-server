package oauthRepository

import (
	articleDomain "github.com/diki-haryadi/go-micro-template/internal/oauth/domain"
	"github.com/diki-haryadi/ztools/postgres"
)

type repository struct {
	postgres *postgres.Postgres
}

func NewRepository(conn *postgres.Postgres) articleDomain.Repository {
	return &repository{postgres: conn}
}
