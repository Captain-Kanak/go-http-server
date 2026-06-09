package main

import (
	"context"
	"fmt"
	"go-http-server/db"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	db.ConnectDb()

	defer db.Db.Close(context.Background())

	server()

	fmt.Println("Program is closed")
}
