package users

import (
	"net/http"

	"github.com/Julio-Cesar07/gobid/internal/api/utils"
	errorsapi "github.com/Julio-Cesar07/gobid/internal/services/errors"
)

func (us *UserHandler) handleLogoutUser(w http.ResponseWriter, r *http.Request) {
	if err := us.Sessions.RenewToken(r.Context()); err != nil {
		utils.EncodeJson(w, utils.Response{Error: errorsapi.ErrSomethingWentWrong.Error()}, http.StatusUnprocessableEntity)
		return
	}

	us.Sessions.Remove(r.Context(), "AuthenticatedUserId")
	utils.EncodeJson(w, utils.Response{Data: "logged out successfully"}, http.StatusOK)
}
