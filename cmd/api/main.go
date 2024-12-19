package main

import (
	"log"

	"github.com/joho/godotenv"

	"gihub.com/kengkeng852/SocialWebsiteGo/internal/env"
	"gihub.com/kengkeng852/SocialWebsiteGo/internal/store"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	store := store.NewStorage(nil)

	app := application{
		config: cfg,
		store: store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
