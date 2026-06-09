package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

// * Connect to DB
func connectDb() {
	var err error

	DATABASE_URL := os.Getenv("DATABASE_URL")

	db, err = pgx.Connect(context.Background(), DATABASE_URL)

	if err != nil {
		panic(err)
	}

	fmt.Println("Database connected successfully!")
}
