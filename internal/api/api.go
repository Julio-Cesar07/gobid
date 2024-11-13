package api

import (
	"net/http"
	"sync"
	"time"

	"github.com/Julio-Cesar07/gobid/internal/api/http/products"
	"github.com/Julio-Cesar07/gobid/internal/api/http/users"
	auctions_service "github.com/Julio-Cesar07/gobid/internal/services/auctions"
	bid_service "github.com/Julio-Cesar07/gobid/internal/services/bids"
	product_service "github.com/Julio-Cesar07/gobid/internal/services/products"
	user_service "github.com/Julio-Cesar07/gobid/internal/services/users"
	"github.com/alexedwards/scs/pgxstore"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Api struct {
	Router *chi.Mux

	UserHandler    *users.UserHandler
	ProductHandler *products.ProductHandler

	Sessions *scs.SessionManager
}

func CreateApi(pool *pgxpool.Pool) Api {
	s := scs.New()
	s.Store = pgxstore.New(pool)
	s.Lifetime = 24 * time.Hour
	s.Cookie.HttpOnly = true
	s.Cookie.SameSite = http.SameSiteLaxMode

	wsUpgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	return Api{
		Router:   chi.NewMux(),
		Sessions: s,
		UserHandler: &users.UserHandler{
			Service:  user_service.NewUserService(pool),
			Sessions: s,
		},
		ProductHandler: &products.ProductHandler{
			ProductsService: product_service.NewProductService(pool),
			Sessions:        s,
			BidsService:     bid_service.NewBidsService(pool),
			WsUpgrager:      wsUpgrader,
			AuctionLobby: auctions_service.AuctionLobby{
				Rooms: make(map[uuid.UUID]*auctions_service.AuctionRoom),
				Mutex: sync.Mutex{},
			},
		},
	}
}
