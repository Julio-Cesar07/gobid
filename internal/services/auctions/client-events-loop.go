package auctions

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/gorilla/websocket"
)

const (
	maxMessageSize = 512
	readDeadline   = 60 * time.Second
	writeWait      = 10 * time.Second
	pingPeriod     = (readDeadline * 9) / 10
)

func (c *Client) ReadEventLoop() {
	defer func() {
		c.Room.Unregister <- c
		c.Conn.Close()
	}()

	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(readDeadline))
	c.Conn.SetPongHandler(func(appData string) error {
		c.Conn.SetReadDeadline(time.Now().Add(readDeadline))
		return nil
	})

	for {
		var m Message
		m.UserId = c.UserId
		m.Kind = -1
		if err := c.Conn.ReadJSON(&m); err != nil || m.Amount == 0 || m.Kind == -1 {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				slog.Error("Unexpected Close error", "error", err)
				return
			}

			c.Room.Broadcast <- Message{
				Kind:    InvalidJson,
				Message: "invalid json",
				UserId:  c.UserId,
			}
			continue
		}
		fmt.Println(m)

		c.Room.Broadcast <- m

	}
}

func (c *Client) WriteEventLoop() {
	ticker := time.NewTicker(pingPeriod)

	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.Conn.WriteJSON(Message{
					Kind:    CloseConnection,
					Message: "closing websocket conn",
				})
				return
			}

			if message.Kind == AuctionFinished {
				close(c.Send)
				return
			}
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))

			err := c.Conn.WriteJSON(message)
			if err != nil {
				c.Room.Unregister <- c
				return
			}

		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				slog.Error("Unexpected write error", "error", err)
				return
			}
		}
	}
}
