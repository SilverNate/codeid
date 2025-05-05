package main

import (
	"codeid/config"
	"codeid/database"
	"codeid/router"
	"log"
)

func main() {
	cfg := config.Load()

	db := database.InitDB(cfg.DBDSN)

	r := router.InitRouter(db)
	if err := r.Run(":8082"); err != nil {
		log.Fatal("Failed to start server:", err)
	}

	r.Run(":8082")
}
