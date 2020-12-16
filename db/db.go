package db

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/anihouse/bot"
	"github.com/anihouse/bot/config"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/log/logrusadapter"
	"github.com/sirupsen/logrus"
)

var (
	pgxconn *pgx.Conn
	logger  *logrus.Logger
)

// Tables
var (
	Clubs clubs
)

func Init() {
	fmt.Println("Database connection...")

	connCfg, err := pgx.ParseURI(config.DB.Connection)
	if err != nil {
		log.Fatal(err)
	}

	logger, err = bot.Logger(config.DB.Log)
	if err != nil {
		if !errors.Is(err, bot.ErrNoLogger) {
			log.Fatal(err)
		}
	}

	if logger != nil {
		connCfg.Logger = logrusadapter.NewLogger(logger)
	}

	pool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig:     connCfg,
		MaxConnections: 20,
		AcquireTimeout: 30 * time.Second,
	})

	if err != nil {
		log.Fatal(err)
	}

	pgxconn, err = pool.Acquire()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected")
}
