package products

import (
	"context"
	"errors"
	"log/slog"
	"strconv"
	"time"

	errorsapi "github.com/Julio-Cesar07/gobid/internal/services/errors"
	"github.com/Julio-Cesar07/gobid/internal/store/pgstore"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateProductReq struct {
	Selled_id    uuid.UUID
	Product_name string
	Description  string
	Baseprice    float64
	Auction_end  time.Time
}

func (ps *ProductService) CreateProduct(ctx context.Context, params CreateProductReq) (uuid.UUID, error) {
	var basepricePgType pgtype.Numeric
	basepriceString := strconv.FormatFloat(params.Baseprice, 'f', -1, 64)

	if err := basepricePgType.Scan(basepriceString); err != nil {
		slog.Error("failed to convert float64 to pgtype.Numeric", "error", err)
		return uuid.UUID{}, err
	}

	args := pgstore.CreateProductParams{
		SellerID:    params.Selled_id,
		ProductName: params.Product_name,
		Description: pgtype.Text{
			String: params.Description,
			Valid:  params.Description != "",
		},
		AuctionEnd: params.Auction_end,
		Baseprice:  basepricePgType,
	}

	id, err := ps.queries.CreateProduct(ctx, args)

	if err != nil {
		var pgError *pgconn.PgError

		if errors.As(err, &pgError) && pgError.Code == "23503" {
			return uuid.UUID{}, errorsapi.ErrSellerNotExist
		}
		slog.Error("failed to create product", "error", err)
		return uuid.UUID{}, err
	}

	return id, nil
}
