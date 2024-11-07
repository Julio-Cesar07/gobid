package users

import (
	user_services "github.com/Julio-Cesar07/gobid/internal/services/users"
	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	Service user_services.UserService
}

func (uh *UserHandler) BindUserRoutes(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Post("/signup", uh.handleSignupUser)
		r.Post("/login", uh.handleLoginUser)
		r.Post("/logout", uh.handleLogoutUser)
	})
}
