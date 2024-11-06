package api

import (
	"net/http"

	"github.com/Julio-Cesar07/gobid/internal/api/utils"
	"github.com/go-chi/chi/v5"
)

func (api *Api) BindRoutes() {
	api.Router.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
				utils.EncodeJson(w, utils.Response{Data: "Server Running"}, 200)
			})
			api.UserHandler.BindUserRoutes(r)
		})
	})
}
