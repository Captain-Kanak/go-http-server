package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var Db *pgx.Conn

func ConnectDb() {
	var err error

	DATABASE_URL := os.Getenv("DATABASE_URL")

	Db, err = pgx.Connect(context.Background(), DATABASE_URL)

	if err != nil {
		panic(err)
	}

	fmt.Println("Database connected successfully!")
}
