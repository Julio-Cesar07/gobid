package products

import (
	"errors"
	"net/http"

	"github.com/Julio-Cesar07/gobid/internal/api/utils"
	"github.com/Julio-Cesar07/gobid/internal/services/auctions"
	errorsapi "github.com/Julio-Cesar07/gobid/internal/services/errors"
	"github.com/Julio-Cesar07/gobid/internal/services/products"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (ph *ProductHandler) handleSubscribeUserToAuction(w http.ResponseWriter, r *http.Request) {
	rawProductId := chi.URLParam(r, "product_id")

	productId, err := uuid.Parse(rawProductId)

	if err != nil {
		utils.EncodeJson(w, utils.Response{Error: "invalid product id - must be a valid uuid"}, http.StatusBadRequest)
		return
	}

	product, err := ph.ProductsService.GetProductByd(r.Context(), products.GetProductByIdReq{
		ProductID: productId,
	})

	if err != nil {
		if errors.Is(err, errorsapi.ErrProductNotExist) {
			utils.EncodeJson(w, utils.Response{Error: err.Error()}, http.StatusBadRequest)
			return
		}
		utils.EncodeJson(w, utils.Response{Error: errorsapi.ErrSomethingWentWrong.Error()}, http.StatusInternalServerError)
		return
	}

	userId, ok := ph.Sessions.Get(r.Context(), "AuthenticatedUserId").(uuid.UUID)

	if !ok {
		utils.EncodeJson(w, utils.Response{Error: errorsapi.ErrInvalidCredentials.Error()}, http.StatusUnauthorized)
		return
	}

	ph.AuctionLobby.Lock()
	room, found := ph.AuctionLobby.Rooms[product.ID]
	ph.AuctionLobby.Unlock()
	if !found {
		utils.EncodeJson(w, utils.Response{Error: "the auction has ended."}, http.StatusBadRequest)
		return
	}

	conn, err := ph.WsUpgrager.Upgrade(w, r, nil)
	if err != nil {
		utils.EncodeJson(w, utils.Response{Error: "could not upgrade connection to a websocket protocol"}, http.StatusInternalServerError)
		return
	}

	client := auctions.NewClient(room, conn, userId)

	room.Register <- client
	go client.ReadEventLoop()
	go client.WriteEventLoop()

}
