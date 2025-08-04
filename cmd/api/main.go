package main

import (
	"BackendEngineeringGo/internal/env"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	envAddr := os.Getenv("ADDR")

	cfg := config{
		addr: env.GetString("ADDR", envAddr),
	}

	app := &application{
		config: cfg,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
