package api

import (
	"github.com/Julio-Cesar07/gobid/internal/api/users"
	user_service "github.com/Julio-Cesar07/gobid/internal/services/users"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Api struct {
	Router *chi.Mux

	UserHandler *users.UserHandler
}

func CreateApi(pool *pgxpool.Pool) Api {
	return Api{
		Router: chi.NewMux(),
		UserHandler: &users.UserHandler{
			Service: user_service.NewUserService(pool),
		},
	}
}
