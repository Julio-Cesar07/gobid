package users

import (
	"net/http"

	"github.com/Julio-Cesar07/gobid/internal/api/auth"
	user_services "github.com/Julio-Cesar07/gobid/internal/services/users"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	Service user_services.UserService

	Sessions *scs.SessionManager
}

func (uh *UserHandler) BindUserRoutes(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Post("/signup", uh.handleSignupUser)
		r.Post("/sessions", uh.handleAuthenticate)

		r.Group(func(r chi.Router) {
			r.Use(func(h http.Handler) http.Handler {
				return auth.AuthMiddleware(h, &auth.Handler{Sessions: uh.Sessions})
			})

			r.Post("/logout", uh.handleLogoutUser)
		})

	})
}
