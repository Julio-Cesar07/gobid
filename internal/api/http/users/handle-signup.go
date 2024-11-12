package users

import (
	"errors"
	"net/http"

	"github.com/Julio-Cesar07/gobid/internal/api/dtos"
	"github.com/Julio-Cesar07/gobid/internal/api/utils"
	errorsapi "github.com/Julio-Cesar07/gobid/internal/services/errors"
	"github.com/Julio-Cesar07/gobid/internal/services/users"
)

func (uh *UserHandler) handleSignupUser(w http.ResponseWriter, r *http.Request) {
	data, problems, err := utils.DecodeValidJson[dtos.CreateUserDto](r, dtos.CreateUserDto{})

	if err != nil {
		if problems == nil {
			utils.EncodeJson(w, utils.Response{Error: err.Error()}, http.StatusBadRequest)
			return
		}
		utils.EncodeJson(w, utils.Response{Error: problems}, http.StatusUnprocessableEntity)
		return
	}

	id, err := uh.Service.CreateUser(r.Context(), users.CreateUserReq{
		Username: data.Username,
		Email:    data.Email,
		Password: data.Password,
		Bio:      data.Bio,
	})

	if err != nil {
		if errors.Is(err, errorsapi.ErrDuplicatedEmailOrUsername) {
			utils.EncodeJson(w, utils.Response{Error: errorsapi.ErrDuplicatedEmailOrUsername.Error()}, http.StatusUnprocessableEntity)
			return
		}
		utils.EncodeJson(w, utils.Response{Error: errorsapi.ErrSomethingWentWrong.Error()}, http.StatusUnprocessableEntity)
		return
	}

	type response struct {
		UserId string `json:"user_id"`
	}

	utils.EncodeJson(w, utils.Response{Data: response{UserId: id.String()}}, http.StatusCreated)
}
