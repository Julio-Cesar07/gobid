package products

import (
	"net/http"

	"github.com/Julio-Cesar07/gobid/internal/api/auth"
	product_service "github.com/Julio-Cesar07/gobid/internal/services/products"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	Service product_service.ProductService

	Sessions *scs.SessionManager
}

func (ph *ProductHandler) BindProductsRoutes(r chi.Router) {
	r.Route("/products", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(func(h http.Handler) http.Handler {
				return auth.AuthMiddleware(h, &auth.Handler{Sessions: ph.Sessions})
			})

			r.Post("/", ph.handleCreateProduct)
		})
	})
}
