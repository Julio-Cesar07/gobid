package users

import (
	"context"
	"errors"
	"log/slog"

	errorsapi "github.com/Julio-Cesar07/gobid/internal/services/errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthenticateReq struct {
	Email    string
	Password string
}

type AuthenticateRes struct {
	Id uuid.UUID
}

func (us *UserService) Authenticate(ctx context.Context, data AuthenticateReq) (AuthenticateRes, error) {
	user, err := us.queries.GetUserByEmail(ctx, data.Email)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return AuthenticateRes{}, errorsapi.ErrInvalidCredentials
		}
		slog.Error("failed to get user by email", "error", err)
		return AuthenticateRes{Id: uuid.UUID{}}, err
	}

	if err = bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(data.Password)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return AuthenticateRes{Id: uuid.UUID{}}, errorsapi.ErrInvalidCredentials
		}
		slog.Error("failed to compare hash and password", "error", err)
		return AuthenticateRes{Id: uuid.UUID{}}, err
	}

	return AuthenticateRes{user.ID}, nil
}
