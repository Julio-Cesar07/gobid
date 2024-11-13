package products

import (
	"net/http"

	"github.com/Julio-Cesar07/gobid/internal/api/auth"
	auctions_service "github.com/Julio-Cesar07/gobid/internal/services/auctions"
	bids_service "github.com/Julio-Cesar07/gobid/internal/services/bids"
	products_service "github.com/Julio-Cesar07/gobid/internal/services/products"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
)

type ProductHandler struct {
	ProductsService products_service.ProductService
	BidsService     bids_service.BidsService
	AuctionLobby    auctions_service.AuctionLobby

	Sessions *scs.SessionManager

	WsUpgrager websocket.Upgrader
}

func (ph *ProductHandler) BindProductsRoutes(r chi.Router) {
	r.Route("/products", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(func(h http.Handler) http.Handler {
				return auth.AuthMiddleware(h, &auth.Handler{Sessions: ph.Sessions})
			})

			r.Post("/", ph.handleCreateProduct)

			r.Get("/ws/subscribe/{product_id}", ph.handleSubscribeUserToAuction)
		})
	})
}
