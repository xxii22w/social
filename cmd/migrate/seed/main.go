package main

import (
	"log"
	"os"

	"github.com/xxii22w/social/internal/db"
	"github.com/xxii22w/social/internal/store"
)

func main() {
	addr := os.Getenv("DB_ADDR")
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	store := store.NewStorage(conn)
	db.Seed(store,conn)
}
