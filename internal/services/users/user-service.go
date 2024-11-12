package users

import (
	"github.com/Julio-Cesar07/gobid/internal/store/pgstore"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserService struct {
	pool    *pgxpool.Pool
	queries *pgstore.Queries
}

func NewUserService(pool *pgxpool.Pool) UserService {
	return UserService{
		pool:    pool,
		queries: pgstore.New(pool),
	}
}
