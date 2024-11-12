package users

import (
	"errors"
	"net/http"

	"github.com/Julio-Cesar07/gobid/internal/api/dtos"
	"github.com/Julio-Cesar07/gobid/internal/api/utils"
	errorsapi "github.com/Julio-Cesar07/gobid/internal/services/errors"
	"github.com/Julio-Cesar07/gobid/internal/services/users"
)

func (uh *UserHandler) handleAuthenticate(w http.ResponseWriter, r *http.Request) {
	data, problems, err := utils.DecodeValidJson[dtos.AuthenticateDto](r, dtos.AuthenticateDto{})

	if err != nil {
		if problems == nil {
			utils.EncodeJson(w, utils.Response{Error: err.Error()}, http.StatusBadRequest)
			return
		}
		utils.EncodeJson(w, utils.Response{Error: problems}, http.StatusUnprocessableEntity)
		return
	}

	result, err := uh.Service.Authenticate(r.Context(), users.AuthenticateReq{Email: data.Email, Password: data.Password})

	if err != nil {
		if errors.Is(err, errorsapi.ErrInvalidCredentials) {
			utils.EncodeJson(w, utils.Response{Error: errorsapi.ErrInvalidCredentials.Error()}, http.StatusBadRequest)
			return
		}
		utils.EncodeJson(w, utils.Response{Error: errorsapi.ErrSomethingWentWrong.Error()}, http.StatusUnprocessableEntity)
		return
	}

	if err = uh.Sessions.RenewToken(r.Context()); err != nil {
		utils.EncodeJson(w, utils.Response{Error: errorsapi.ErrSomethingWentWrong.Error()}, http.StatusUnprocessableEntity)
		return
	}

	uh.Sessions.Put(r.Context(), "AuthenticatedUserId", result.Id)

	utils.EncodeJson(w, utils.Response{Data: "loggend in, successfully"}, http.StatusOK)
}
