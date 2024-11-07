package users

import (
	"context"
	"errors"
	"log/slog"

	"github.com/Julio-Cesar07/gobid/internal/store/pgstore"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserReq struct {
	Username string
	Email    string
	Password string
	Bio      string
}

func (us *UserService) CreateUser(ctx context.Context, data CreateUserReq) (uuid.UUID, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), 12)

	if err != nil {
		slog.Error("failed to generate uuid", "error", err)
		return uuid.UUID{}, err
	}

	args := pgstore.CreateUserParams{
		Username:     data.Username,
		Email:        data.Email,
		PasswordHash: hash,
		Bio: pgtype.Text{
			String: data.Bio,
			Valid:  data.Bio != "",
		},
	}

	id, err := us.queries.CreateUser(ctx, args)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return uuid.UUID{}, ErrDuplicatedEmailOrUsername
		}
		slog.Error("failed to create user in database")
		return uuid.UUID{}, err
	}

	return id, nil
}
