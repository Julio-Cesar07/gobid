package products

import (
	"errors"
	"net/http"

	"github.com/Julio-Cesar07/gobid/internal/api/dtos"
	"github.com/Julio-Cesar07/gobid/internal/api/utils"
	errorsapi "github.com/Julio-Cesar07/gobid/internal/services/errors"
	"github.com/Julio-Cesar07/gobid/internal/services/products"
	"github.com/google/uuid"
)

func (ph *ProductHandler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	userId, ok := ph.Sessions.Get(r.Context(), "AuthenticatedUserId").(uuid.UUID)

	if !ok {
		utils.EncodeJson(w, utils.Response{Error: errorsapi.ErrInvalidCredentials.Error()}, http.StatusUnauthorized)
		return
	}

	data, problems, err := utils.DecodeValidJson[dtos.CreateProductDto](r, dtos.CreateProductDto{
		SellerID: userId.String(),
	})

	if err != nil {
		if problems == nil {
			utils.EncodeJson(w, utils.Response{Error: err.Error()}, http.StatusBadRequest)
			return
		}
		utils.EncodeJson(w, utils.Response{Error: problems}, http.StatusUnprocessableEntity)
		return
	}

	id, err := ph.Service.CreateProduct(r.Context(), products.CreateProductReq{
		Selled_id:    userId,
		Product_name: data.ProductName,
		Description:  data.Description,
		Baseprice:    data.Baseprice,
		Auction_end:  data.AuctionEnd,
	})

	if err != nil {
		if errors.Is(err, errorsapi.ErrSellerNotExist) {
			utils.EncodeJson(w, utils.Response{Error: errorsapi.ErrSellerNotExist.Error()}, http.StatusBadRequest)
			return
		}
		utils.EncodeJson(w, utils.Response{Error: errorsapi.ErrSomethingWentWrong.Error()}, http.StatusInternalServerError)
		return
	}

	type response struct {
		ProductId string `json:"product_id"`
	}

	utils.EncodeJson(w, utils.Response{Data: response{ProductId: id.String()}}, http.StatusCreated)
}
