package api

import (
	"context"
	"encoding/gob"
	"fmt"
	"net/http"
	"os"

	"github.com/Julio-Cesar07/gobid/internal/api"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func init() {
	gob.Register(uuid.UUID{})
}

func Main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
		os.Getenv("GOBID_DATABASE_USER"),
		os.Getenv("GOBID_DATABASE_PASSWORD"),
		os.Getenv("GOBID_DATABASE_HOST"),
		os.Getenv("GOBID_DATABASE_PORT"),
		os.Getenv("GOBID_DATABASE_NAME"),
	))

	if err != nil {
		panic(err)
	}

	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		panic(err)
	}

	api := api.CreateApi(pool)
	api.BindRoutes()

	fmt.Println("Starting server on port :8080")
	if err := http.ListenAndServe(":8080", api.Router); err != nil {
		panic(err)
	}
}
