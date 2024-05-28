package main

import (
	"log"
	"os"

	"github.com/higorjsilva/goapi/cmd/api"
	"github.com/joho/godotenv"
)

func main()  {
	godotenv.Load()
	server := api.NewAPIServer(":" +  os.Getenv("PORT"), nil)
	
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}