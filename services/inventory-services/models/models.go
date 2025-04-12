package models

import (
	"errors"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	ErrInvalidDBType = errors.New("invalid db type")
)

type ModelService struct {
	db *gorm.DB
}

func New(dbType, dsn string) (*ModelService, error) {
	var (
		db  *gorm.DB
		err error
	)
	switch dbType {
	case "postgres":
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	default:
		return nil, ErrInvalidDBType
	}

	if err != nil {
		return nil, err
	}
	return &ModelService{
		db: db,
	}, nil
}

func (s *ModelService) Migrate() {
	s.db.AutoMigrate(&InventoryProduct{})

}
