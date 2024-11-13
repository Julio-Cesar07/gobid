package products

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/Julio-Cesar07/gobid/internal/api/dtos"
	"github.com/Julio-Cesar07/gobid/internal/api/utils"
	"github.com/Julio-Cesar07/gobid/internal/services/bids"
	errorsapi "github.com/Julio-Cesar07/gobid/internal/services/errors"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (ph *ProductHandler) handlePlaceBid(w http.ResponseWriter, r *http.Request) {
	bidderUUID, ok := ph.Sessions.Get(r.Context(), "AuthenticatedUserId").(uuid.UUID)

	if !ok {
		utils.EncodeJson(w, utils.Response{Error: errorsapi.ErrInvalidCredentials.Error()}, http.StatusUnauthorized)
		return
	}

	rawProductId := chi.URLParam(r, "product_id")

	data, problems, err := utils.DecodeValidJson[dtos.PlaceBidDto](r, dtos.PlaceBidDto{
		ProductID: rawProductId,
		BidderID:  bidderUUID.String(),
	})

	if err != nil {
		if len(problems) == 0 {
			utils.EncodeJson(w, utils.Response{Error: err.Error()}, http.StatusBadRequest)
			return
		}
		utils.EncodeJson(w, utils.Response{Error: problems}, http.StatusUnprocessableEntity)
		return
	}

	productUUID, err := uuid.Parse(data.ProductID)

	if err != nil {
		utils.EncodeJson(w, utils.Response{Error: "invalid product id - must be a valid uuid"}, http.StatusBadRequest)
		return
	}

	bid, err := ph.BidsService.PlaceBid(r.Context(), bids.PlaceBidReq{
		ProductID: productUUID,
		BidderID:  bidderUUID,
		BidAmount: data.BidAmount,
	})

	if err != nil {
		if errors.Is(err, errorsapi.ErrProductNotExist) ||
			errors.Is(err, errorsapi.ErrBidderNotExist) ||
			errors.Is(err, errorsapi.ErrBidLowerThanTheLast) ||
			errors.Is(err, errorsapi.ErrBidLowerThanProductBaseprice) {
			utils.EncodeJson(w, utils.Response{Error: err.Error()}, http.StatusBadRequest)
			return
		}
		slog.Error("failed to create bid", "error", err)
		utils.EncodeJson(w, utils.Response{Error: errorsapi.ErrSomethingWentWrong.Error()}, http.StatusInternalServerError)
		return
	}

	type response struct {
		BidId string `json:"bid_id"`
	}

	utils.EncodeJson(w, utils.Response{Data: response{BidId: bid.ID.String()}}, http.StatusCreated)
}
