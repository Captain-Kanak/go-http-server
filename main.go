package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	connectDb()

	defer db.Close(context.Background())

	server()

	fmt.Println("Program is closed")
}
