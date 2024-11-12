package dtos

import (
	"context"
	"time"
)

type CreateProductDto struct {
	SellerID    string    `json:"seller_id"`
	ProductName string    `json:"product_name"`
	Description string    `json:"description"`
	Baseprice   float64   `json:"baseprice"`
	AuctionEnd  time.Time `json:"auction_end"`
}

func (dto CreateProductDto) Valid(ctx context.Context) Evaluator {
	var eval Evaluator

	eval.CheckField(NotBlank(dto.SellerID), "seller_id", "this field cannot be empty")
	eval.CheckField(NotBlank(dto.ProductName), "product_name", "this field cannot be empty")
	eval.CheckField(MaxChars(dto.Description, 255), "description", "this field must have a maximum length of 255")
	eval.CheckField(dto.Baseprice > 0, "baseprice", "this field must be grater than 0")
	eval.CheckField(Float2Decimals(dto.Baseprice), "baseprice", "this field must have 2 decimal places")
	eval.CheckField(time.Until(dto.AuctionEnd) >= minAuctionDuration, "auction_end", "must be at least two hours duration")

	return eval
}
