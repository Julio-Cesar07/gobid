package auctions

import (
	"log/slog"
)

func (r *AuctionRoom) Run() {
	slog.Info("Auction has begun", "auctionId", r.ProductId)
	defer func() {
		close(r.Broadcast)
		close(r.Register)
		close(r.Unregister)
	}()

	for {
		select {
		case client := <-r.Register:
			r.registerClient(client)

		case client := <-r.Unregister:
			r.unregisterClient(client)

		case message := <-r.Broadcast:
			r.broadcastMessage(message)

		case <-r.Context.Done():
			slog.Info("Auction has ended", "auctionId", r.ProductId)
			for _, client := range r.Clients {
				client.Send <- Message{Kind: AuctionFinished, Message: "auction has been finished"}
			}
			return
		}
	}
}
