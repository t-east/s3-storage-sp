package http

import (
	"sp/src/domains/entities"
	"sp/src/interfaces/contracts"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func LoadTestDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&entities.User{})
	return db, nil
}

func LoadTestParam() (*contracts.Param, error) {
	return &contracts.Param{}, nil
}
