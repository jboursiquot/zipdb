package zipdb

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DB struct {
	conn *gorm.DB
}

func NewDB(path string) (*DB, error) {
	conn, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := conn.AutoMigrate(&Location{}); err != nil {
		return nil, err
	}
	return &DB{conn: conn}, nil
}

func (db *DB) Seed(locations map[string]Location) error {
	for _, l := range locations {
		if err := db.conn.Clauses(clause.OnConflict{UpdateAll: true}).Create(&l).Error; err != nil {
			return err
		}
	}
	return nil
}

func (db *DB) Find(zip string) (*Location, error) {
	var loc Location
	if err := db.conn.First(&loc, zip).Error; err != nil {
		return nil, err
	}
	return &loc, nil
}
