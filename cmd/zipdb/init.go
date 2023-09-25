package main

import (
	"log/slog"
	"os"

	"github.com/jboursiquot/zipdb"
	"github.com/joeshaw/envdecode"
)

type config struct {
	Host   string `env:"HOST,default=127.0.0.1"`
	Port   string `env:"PORT,default=8080"`
	DBFile string `env:"DB_FILE,default=data/US.db"`
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
