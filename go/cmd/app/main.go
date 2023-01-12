package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {

	fmt.Println("Hello world")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	// client.Db.ConnectDb()
}
