package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/xxii22w/social/internal/db"
	"github.com/xxii22w/social/internal/store"
	"go.uber.org/zap"
)

const version = "0.0.1"

//	@title			Social
//	@description	API for Social
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@BasePath					/v1
//
//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	maxOpenConns, maxIdleConns := LoadEnvFromStringToInt64()

	cfg := config{
		addr:   os.Getenv("Addr"),
		apiURL: os.Getenv("EXTERNAL_URL"),
		db: dbConfig{
			addr:         os.Getenv("DB_ADDR"),
			maxOpenConns: maxOpenConns,
			maxIdleConns: maxIdleConns,
			maxIdleTime:  os.Getenv("DB_MAX_IDLE_TIME"),
		},
		env: os.Getenv("ENV"),
		mail: mailConfig{
			exp: time.Hour * 24 * 3,
		},
	}

	// Logger
	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	db, err := db.New(
		cfg.db.addr,
		int(cfg.db.maxOpenConns),
		int(cfg.db.maxIdleConns),
		cfg.db.maxIdleTime,
	)
	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()
	logger.Info("database connection pool established")

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
		logger: logger,
	}

	mux := app.mount()
	logger.Fatal(app.run(mux))
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
