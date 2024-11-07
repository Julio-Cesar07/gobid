package users

import (
	"errors"
	"net/http"

	"github.com/Julio-Cesar07/gobid/internal/api/dtos"
	"github.com/Julio-Cesar07/gobid/internal/api/utils"
	"github.com/Julio-Cesar07/gobid/internal/services/users"
)

func (uh *UserHandler) handleSignupUser(w http.ResponseWriter, r *http.Request) {
	data, problems, err := utils.DecodeValidJson[dtos.CreateUserDto](r)

	if err != nil {
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
		if errors.Is(err, users.ErrDuplicatedEmailOrUsername) {
			utils.EncodeJson(w, utils.Response{Error: users.ErrDuplicatedEmailOrUsername.Error()}, http.StatusUnprocessableEntity)
			return
		}
		utils.EncodeJson(w, utils.Response{Error: "something went wrong"}, http.StatusUnprocessableEntity)
		return
	}

	type response struct {
		UserId string `json:"user_id"`
	}

	utils.EncodeJson(w, utils.Response{Data: response{UserId: id.String()}}, http.StatusCreated)
}

func (uh *UserHandler) handleLogoutUser(w http.ResponseWriter, r *http.Request) {

}
