package users

import (
	"context"
	"errors"

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

func (us *UserService) CreateUser(ctx context.Context, user CreateUserReq) (uuid.UUID, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)

	if err != nil {
		return uuid.UUID{}, err
	}

	args := pgstore.CreateUserParams{
		Username:     user.Username,
		Email:        user.Email,
		PasswordHash: hash,
		Bio: pgtype.Text{
			String: user.Bio,
			Valid:  user.Bio != "",
		},
	}

	id, err := us.queries.CreateUser(ctx, args)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return uuid.UUID{}, ErrDuplicatedEmailOrUsername
		}
		return uuid.UUID{}, err
	}

	return id, nil
}
