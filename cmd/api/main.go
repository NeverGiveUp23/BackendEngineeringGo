package main

import (
	"BackendEngineeringGo/internal/env"
	store2 "BackendEngineeringGo/internal/store"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	ENVADDR  = os.Getenv("ADDR")
	GODOTENV = godotenv.Load()
)

func main() {
	err := GODOTENV
	if err != nil {
		log.Fatal("error loading .env file")
	}

	cfg := config{
		addr: env.GetString("ADDR", ENVADDR),
	}

	store := store2.NewStorage(nil)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
