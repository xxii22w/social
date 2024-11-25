package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/xxii22w/social/internal/db"
	"github.com/xxii22w/social/internal/store"
)

const version = "0.0.1"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	maxOpenConns,maxIdleConns := LoadEnvFromStringToInt64()

	cfg := config{
		addr: os.Getenv("Addr"),
		db: dbConfig{
			addr:         os.Getenv("DB_ADDR"),
			maxOpenConns: maxOpenConns,
			maxIdleConns: maxIdleConns,
			maxIdleTime: os.Getenv("DB_MAX_IDLE_TIME"),
		},
		env: os.Getenv("ENV"),
	}

	db,err := db.New(
		cfg.db.addr,
		int(cfg.db.maxOpenConns),
		int(cfg.db.maxIdleConns),
		cfg.db.maxIdleTime,
	)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Println("database connection pool established")

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	log.Printf("server has started at %s", app.config.addr)

	mux := app.mount()
	log.Fatal(app.run(mux))
}

func LoadEnvFromStringToInt64() (int64, int64) {
	maxOpenConns, err := strconv.ParseInt(os.Getenv("DB_MAX_OPEN_CONNS"), 10, 64)
	if err != nil {
		fmt.Errorf(".env enviment loading error %s", err)
	}
	maxIdleConns, err := strconv.ParseInt(os.Getenv("DB_MAX_IDLE_CONNS"), 10, 64)
	if err != nil {
		fmt.Errorf(".env enviment loading error: %s", err)
	}
	return maxOpenConns, maxIdleConns
}
