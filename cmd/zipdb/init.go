package main

import (
	"log/slog"
	"os"

	"github.com/jboursiquot/zipdb"
	"github.com/joeshaw/envdecode"
)

type config struct {
	Host     string `env:"HOST,default=127.0.0.1"`
	Port     string `env:"PORT,default=8080"`
	DataFile string `env:"DATA_FILE,default=data/US.txt"`
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
