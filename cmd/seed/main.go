package main

import (
	"os"

	"github.com/jboursiquot/zipdb"
)

func main() {
	locations, err := zipdb.LoadLocations(cfg.DataFile)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	db, err := zipdb.NewDB(cfg.DBFile)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	if err := db.Seed(locations); err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	log.Info("Done")
}
