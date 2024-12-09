package main

import (
	"github.com/ViniciusDSLima/AuthSystem/cmd/api"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("Erro ao carregar .env")
	}

	server := api.NewApiServer(os.Getenv("PORT"))

	err := server.Start()

	if err != nil {
		log.Fatal(err)
	}
}
