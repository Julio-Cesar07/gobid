package users

import (
	user_services "github.com/Julio-Cesar07/gobid/internal/services/users"
	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	Service user_services.UserService
}

func (u *UserHandler) BindUserRoutes(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Post("/signup", u.handleSignupUser)
		r.Post("/login", u.handleLoginUser)
		r.Post("/logout", u.handleLogoutUser)
	})
}
