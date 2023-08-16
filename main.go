package main

import (
	"log"
	"github.com/kendax/tic_tac_toe_go_internal/routes"
	"os"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	r := routes.SetupRoutes()
	log.Println("listening on http://localhost:3000")
	r.Run(":" + port)

}
