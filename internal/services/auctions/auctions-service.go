package auctions

import (
	"context"
	"sync"

	bids_service "github.com/Julio-Cesar07/gobid/internal/services/bids"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type MessageKind int

const (
	PlaceBid MessageKind = iota

	// Ok/Success
	SuccessfullyPlacedBid

	// Errors
	FailedToPlaceBid
	FailedToGetProductOrBidder
	InvalidJson

	// Info
	NewBidPlaced
	AuctionFinished
	CloseConnection
)

type Message struct {
	Message string      `json:"message,omitempty"`
	Amount  float64     `json:"amount,omitempty"`
	Kind    MessageKind `json:"kind"`
	UserId  uuid.UUID   `json:"user_id,omitempty"`
}

type AuctionLobby struct {
	sync.Mutex
	Rooms map[uuid.UUID]*AuctionRoom
}

type AuctionRoom struct {
	ProductId  uuid.UUID
	Context    context.Context
	Broadcast  chan Message
	Register   chan *Client
	Unregister chan *Client
	Clients    map[uuid.UUID]*Client

	BidsService bids_service.BidsService
}

func NewAuctionRoom(ctx context.Context, BidsService bids_service.BidsService, productId uuid.UUID) *AuctionRoom {
	return &AuctionRoom{
		ProductId:   productId,
		Context:     ctx,
		BidsService: BidsService,
		Broadcast:   make(chan Message),
		Register:    make(chan *Client),
		Unregister:  make(chan *Client),
		Clients:     make(map[uuid.UUID]*Client),
	}
}

type Client struct {
	Room   *AuctionRoom
	Conn   *websocket.Conn
	Send   chan Message
	UserId uuid.UUID
}

func NewClient(room *AuctionRoom, conn *websocket.Conn, userId uuid.UUID) *Client {
	return &Client{
		Room:   room,
		Conn:   conn,
		Send:   make(chan Message, 512),
		UserId: userId,
	}
}
