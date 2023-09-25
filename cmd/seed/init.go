package main

import (
	"log/slog"
	"os"

	"github.com/jboursiquot/zipdb"
	"github.com/joeshaw/envdecode"
)

type config struct {
	DataFile string `env:"DATA_FILE,default=data/US.txt"`
	DBFile   string `env:"DB_FILE,default=data/US.db"`
}

var (
	log *slog.Logger
	cfg config
)

func init() {
	log = zipdb.DefaultLogger()

	log.Info("Initializing Config...")
	cfg = config{}
	if err := envdecode.Decode(&cfg); err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}
