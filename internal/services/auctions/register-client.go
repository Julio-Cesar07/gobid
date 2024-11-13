package auctions

import "log/slog"

func (r *AuctionRoom) registerClient(c *Client) {
	slog.Info("New User Connected", "Client", c)
	r.Clients[c.UserId] = c
}
