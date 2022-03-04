package main

import (
	"fmt"
	"log"
	"os"

	router "sp/src/drivers/router"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env not found")
	}

	log.Println("Server running...")
	router.Serve(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
