package bids

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strconv"

	errorsapi "github.com/Julio-Cesar07/gobid/internal/services/errors"
	"github.com/Julio-Cesar07/gobid/internal/store/pgstore"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

type PlaceBidReq struct {
	ProductID uuid.UUID `json:"product_id"`
	BidderID  uuid.UUID `json:"bidder_id"`
	BidAmount float64   `json:"bid_amount"`
}

func (bs *BidsService) PlaceBid(ctx context.Context, params PlaceBidReq) (uuid.UUID, error) {
	highestBid, err := bs.queries.GetHighestBidByProductId(ctx, params.ProductID)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			product, err := bs.queries.GetProductById(ctx, params.ProductID)

			if err != nil {
				if errors.Is(err, pgx.ErrNoRows) {
					return uuid.UUID{}, errorsapi.ErrProductNotExist
				}
				slog.Error("failed to get product by id", "error", err)
				return uuid.UUID{}, err
			}

			productBaseprice, err := product.Baseprice.Float64Value()

			if err != nil {
				slog.Error("failed to convert product baseprice to float 64", "error", err)
				return uuid.UUID{}, err
			}

			if productBaseprice.Float64 > params.BidAmount {
				return uuid.UUID{}, errorsapi.ErrBidLowerThanProductBaseprice
			}
		} else {
			slog.Error("failed to get highest bid by product id", "error", err)
			return uuid.UUID{}, err
		}
	} else {
		lastBidAmountFloat, err := highestBid.BidAmount.Float64Value()

		if err != nil {
			slog.Error("failed to convert bid amount to float 64", "error", err)
			return uuid.UUID{}, err
		}

		if lastBidAmountFloat.Float64 >= params.BidAmount {
			return uuid.UUID{}, errorsapi.ErrBidLowerThanTheLast
		}
	}

	var bidAmountPgType pgtype.Numeric

	bidAmountStr := strconv.FormatFloat(params.BidAmount, 'f', -1, 64)

	if err := bidAmountPgType.Scan(bidAmountStr); err != nil {
		return uuid.UUID{}, err
	}

	args := pgstore.CreateBidsParams{
		ProductID: params.ProductID,
		BidderID:  params.BidderID,
		BidAmount: bidAmountPgType,
	}

	id, err := bs.queries.CreateBids(ctx, args)

	if err != nil {
		var pgtypeErr *pgconn.PgError

		if errors.As(err, &pgtypeErr) && pgtypeErr.Code == "23503" {
			fmt.Println(pgtypeErr)
			return uuid.UUID{}, errorsapi.ErrProductNotExist
		}
		slog.Error("failed to create bid", "error", err)
		return uuid.UUID{}, errorsapi.ErrSomethingWentWrong
	}

	return id, nil
}
