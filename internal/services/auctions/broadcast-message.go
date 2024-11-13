package auctions

import (
	"errors"
	"log/slog"

	"github.com/Julio-Cesar07/gobid/internal/services/bids"
	errorsapi "github.com/Julio-Cesar07/gobid/internal/services/errors"
)

func (r *AuctionRoom) broadcastMessage(m Message) {
	slog.Info("New message recieved", "RoomID", r.ProductId, "Message", m.Message, "UserId", m.UserId)

	switch m.Kind {
	case PlaceBid:
		bid, err := r.BidsService.PlaceBid(r.Context, bids.PlaceBidReq{
			ProductID: r.ProductId,
			BidderID:  m.UserId,
			BidAmount: m.Amount,
		})

		if err != nil {
			if errors.Is(err, errorsapi.ErrProductNotExist) ||
				errors.Is(err, errorsapi.ErrBidderNotExist) {
				if client, found := r.Clients[m.UserId]; found {
					client.Send <- Message{Kind: FailedToGetProductOrBidder, Message: err.Error(), UserId: m.UserId}
				}
			} else if errors.Is(err, errorsapi.ErrBidLowerThanTheLast) ||
				errors.Is(err, errorsapi.ErrBidLowerThanProductBaseprice) {
				if client, found := r.Clients[m.UserId]; found {
					client.Send <- Message{Kind: FailedToPlaceBid, Message: err.Error(), UserId: m.UserId}
				}
			}

			return
		}

		if client, found := r.Clients[m.UserId]; found {
			client.Send <- Message{Kind: SuccessfullyPlacedBid, Message: "Your bid was Successfully placed.", UserId: m.UserId}
		} else {
			slog.Info("Client not found in hashmap", "userId", m.UserId)
		}

		for id, client := range r.Clients {
			if id != m.UserId {
				newBidMessage := Message{Kind: NewBidPlaced, Message: "A new bid was placed", Amount: bid.BidAmount, UserId: id}
				client.Send <- newBidMessage
			}
		}
	case InvalidJson:
		if client, found := r.Clients[m.UserId]; found {
			client.Send <- m
		} else {
			slog.Info("Client not found in hashmap", "userId", m.UserId)
		}
	}
}
