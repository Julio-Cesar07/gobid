package dtos

import "context"

type PlaceBidDto struct {
	ProductID string  `json:"product_id"`
	BidderID  string  `json:"bidder_id"`
	BidAmount float64 `json:"bid_amount"`
}

func (dto PlaceBidDto) Valid(ctx context.Context) Evaluator {
	var eval Evaluator

	eval.CheckField(NotBlank(dto.ProductID), "product_id", "this field cannot be empty")
	eval.CheckField(NotBlank(dto.BidderID), "bidder_id", "this field cannot be empty")
	eval.CheckField(dto.BidAmount > 0, "bid_amount", "must be greater than zero")
	eval.CheckField(Float2Decimals(dto.BidAmount), "bid_amount", "this field must have 2 decimal places")

	return eval
}
