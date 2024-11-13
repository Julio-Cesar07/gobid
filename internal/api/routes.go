package api

import (
	"net/http"
	"os"

	"github.com/Julio-Cesar07/gobid/internal/api/auth"
	"github.com/Julio-Cesar07/gobid/internal/api/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/csrf"
)

func (api *Api) BindRoutes() {
	api.Router.Use(middleware.RequestID, middleware.Recoverer, middleware.Logger, api.Sessions.LoadAndSave)

	csrfMiddlewate := csrf.Protect(
		[]byte(os.Getenv("GOBID_CSRF_KEY")),
		csrf.Path("/"),
		csrf.Secure(false), // DEV ONLY
	)

	api.Router.Use(csrfMiddlewate)

	api.Router.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
				utils.EncodeJson(w, utils.Response{Data: "Server Running"}, 200)
			})

			r.Get("/csrftoken", auth.HandleGetCSRFtoken)

			api.UserHandler.BindUserRoutes(r)
			api.ProductHandler.BindProductsRoutes(r)
		})
	})
}
