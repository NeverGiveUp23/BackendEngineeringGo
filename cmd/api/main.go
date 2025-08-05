package main

import (
	"BackendEngineeringGo/internal/db"
	"BackendEngineeringGo/internal/env"
	"BackendEngineeringGo/internal/store"
	"database/sql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	ENVADDR      = os.Getenv("ADDR")
	ENVDBADDR    = os.Getenv("DB_ADDR")
	GODOTENV     = godotenv.Load()
	MAXOPENCONNS = 30
	MAXIDLECONNS = 30
	MAXIDLETIME  = "15m"
)

const version = "0.0.1"

func main() {
	err := GODOTENV
	if err != nil {
		log.Fatal("error loading .env file")
	}

	cfg := config{
		addr: env.GetString("ADDR", ENVADDR),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", ENVDBADDR),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", MAXOPENCONNS),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", MAXIDLECONNS),
			MaxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", MAXIDLETIME),
		},
		env: env.GetString("Env", "development"),
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.MaxIdleTime,
	)
	if err != nil {
		log.Panic(err)
	}

	// close DB
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)

	log.Println("database connect pool established")

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
