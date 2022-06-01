package main

import (
	"github.com/smhdhsn/restaurant-order/internal/config"
	"github.com/smhdhsn/restaurant-order/internal/db"

	log "github.com/smhdhsn/restaurant-order/internal/logger"
)

// main is the application's kernel.
func main() {
	// read configurations.
	conf, err := config.LoadConf()
	if err != nil {
		log.Fatal(err)
	}

	// create a database connection.
	dbConn, err := db.Connect(&conf.DB)
	if err != nil {
		log.Fatal(err)
	}

	// initialize auto migration.
	if err := db.InitMigrations(dbConn); err != nil {
		log.Fatal(err)
	}
}
