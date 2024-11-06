package users

import (
	"errors"
	"net/http"

	"github.com/Julio-Cesar07/gobid/internal/api/dtos"
	"github.com/Julio-Cesar07/gobid/internal/api/utils"
	user_services "github.com/Julio-Cesar07/gobid/internal/services/users"
)

func (u *UserHandler) handleSignupUser(w http.ResponseWriter, r *http.Request) {
	data, problems, err := utils.DecodeValidJson[dtos.CreateUserReq](r)

	if err != nil {
		utils.EncodeJson(w, utils.Response{Error: problems}, http.StatusUnprocessableEntity)
		return
	}

	id, err := u.Service.CreateUser(r.Context(), user_services.CreateUserReq{
		Username: data.Username,
		Email:    data.Email,
		Password: data.Password,
		Bio:      data.Bio,
	})

	if err != nil {
		if errors.Is(err, user_services.ErrDuplicatedEmailOrUsername) {
			utils.EncodeJson(w, utils.Response{Error: "email or username already exists"}, http.StatusUnprocessableEntity)
			return
		}
		utils.EncodeJson(w, utils.Response{Error: "something went wrong"}, http.StatusInternalServerError)
		return
	}

	type response struct {
		UserId string `json:"user_id"`
	}

	utils.EncodeJson(w, utils.Response{Data: response{
		UserId: id.String(),
	}}, http.StatusInternalServerError)
}

func (u *UserHandler) handleLoginUser(w http.ResponseWriter, r *http.Request) {}

func (u *UserHandler) handleLogoutUser(w http.ResponseWriter, r *http.Request) {}
