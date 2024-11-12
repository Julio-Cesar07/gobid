package auth

import (
	"net/http"

	"github.com/Julio-Cesar07/gobid/internal/api/utils"
	"github.com/alexedwards/scs/v2"
	"github.com/gorilla/csrf"
)

type Handler struct {
	Sessions *scs.SessionManager
}

func AuthMiddleware(next http.Handler, h *Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !h.Sessions.Exists(r.Context(), "AuthenticatedUserId") {
			utils.EncodeJson(w, utils.Response{Error: "must be logged in"}, http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func HandleGetCSRFtoken(w http.ResponseWriter, r *http.Request) {
	token := csrf.Token(r)

	type response struct {
		CSRF_Token string `json:"csrf_token"`
	}

	utils.EncodeJson(w, utils.Response{Data: response{CSRF_Token: token}}, http.StatusOK)
}
