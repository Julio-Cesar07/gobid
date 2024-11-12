package api

import (
	"net/http"
	"time"

	"github.com/Julio-Cesar07/gobid/internal/api/http/bids"
	"github.com/Julio-Cesar07/gobid/internal/api/http/products"
	"github.com/Julio-Cesar07/gobid/internal/api/http/users"
	bid_service "github.com/Julio-Cesar07/gobid/internal/services/bids"
	product_service "github.com/Julio-Cesar07/gobid/internal/services/products"
	user_service "github.com/Julio-Cesar07/gobid/internal/services/users"
	"github.com/alexedwards/scs/pgxstore"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Api struct {
	Router *chi.Mux

	UserHandler    *users.UserHandler
	ProductHandler *products.ProductHandler
	BidsHandler    *bids.BidsHandler

	Sessions *scs.SessionManager
}

func CreateApi(pool *pgxpool.Pool) Api {
	s := scs.New()
	s.Store = pgxstore.New(pool)
	s.Lifetime = 24 * time.Hour
	s.Cookie.HttpOnly = true
	s.Cookie.SameSite = http.SameSiteLaxMode

	return Api{
		Router:   chi.NewMux(),
		Sessions: s,
		UserHandler: &users.UserHandler{
			Service:  user_service.NewUserService(pool),
			Sessions: s,
		},
		ProductHandler: &products.ProductHandler{
			Service:  product_service.NewProductService(pool),
			Sessions: s,
		},
		BidsHandler: &bids.BidsHandler{
			Service:  bid_service.NewBidsService(pool),
			Sessions: s,
		},
	}
}
