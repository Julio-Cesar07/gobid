package users

import (
	"net/http"

	"github.com/Julio-Cesar07/gobid/internal/api/dtos"
	"github.com/Julio-Cesar07/gobid/internal/api/utils"
	"github.com/Julio-Cesar07/gobid/internal/services/users"
)

func (uh *UserHandler) handleLoginUser(w http.ResponseWriter, r *http.Request) {
	data, problems, err := utils.DecodeValidJson[dtos.LoginDto](r)

	if err != nil {
		utils.EncodeJson(w, utils.Response{Error: problems}, http.StatusUnprocessableEntity)
		return
	}

	result, err := uh.Service.Authenticate(r.Context(), users.AuthenticateReq{Email: data.Email, Password: data.Password})

	if err != nil {
		utils.EncodeJson(w, utils.Response{Error: "something went wrong"}, http.StatusUnprocessableEntity)
		return
	}

	type response struct {
		Token string `json:"token"`
	}

	utils.EncodeJson(w, utils.Response{Data: response{Token: result.Token}}, http.StatusOK)
}
