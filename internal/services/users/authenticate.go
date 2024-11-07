package users

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type AuthenticateReq struct {
	Email    string
	Password string
}

type AuthenticateRes struct {
	Token string
}

func (us *UserService) Authenticate(ctx context.Context, data AuthenticateReq) (AuthenticateRes, error) {
	_, err := us.queries.GetUserByEmail(ctx, data.Email)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return AuthenticateRes{}, fmt.Errorf("not found")
		}
		return AuthenticateRes{}, fmt.Errorf("not found")
	}

	return AuthenticateRes{Token: "oi"}, nil
}
