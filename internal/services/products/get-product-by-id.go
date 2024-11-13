package products

import (
	"context"
	"errors"
	"log/slog"

	errorsapi "github.com/Julio-Cesar07/gobid/internal/services/errors"
	"github.com/Julio-Cesar07/gobid/internal/store/pgstore"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type GetProductByIdReq struct {
	ProductID uuid.UUID
}

func (ps *ProductService) GetProductByd(ctx context.Context, params GetProductByIdReq) (pgstore.Product, error) {
	product, err := ps.queries.GetProductById(ctx, params.ProductID)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return pgstore.Product{}, errorsapi.ErrProductNotExist
		}
		slog.Error("failed to get product by id", "error", err)
		return pgstore.Product{}, err
	}

	return product, nil
}
