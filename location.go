package zipdb

import (
	"log/slog"

	"gorm.io/gorm"
)

type Location struct {
	gorm.Model
	ID        string `gorm:"primaryKey"`
	Country   string
	Zip       string
	City      string
	StateLong string
	State     string
	County    string
	Lat       float64
	Long      float64
}

func (l *Location) BeforeCreate(tx *gorm.DB) (err error) {
	l.ID = l.Zip
	slog.Info("Adding", "zip", l.Zip)
	return
}
