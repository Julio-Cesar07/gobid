package bids

import (
	"net/http"

	"github.com/Julio-Cesar07/gobid/internal/api/auth"
	bids_service "github.com/Julio-Cesar07/gobid/internal/services/bids"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
)

type BidsHandler struct {
	Service bids_service.BidsService

	Sessions *scs.SessionManager
}

func (bh *BidsHandler) BindBidsRoutes(r chi.Router) {
	r.Route("/bids", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(func(h http.Handler) http.Handler {
				return auth.AuthMiddleware(h, &auth.Handler{Sessions: bh.Sessions})
			})

			r.Post("/{product_id}", bh.handlePlaceBid)
		})
	})
}
