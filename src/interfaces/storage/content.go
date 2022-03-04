package gateways

import (
	"sp/src/domains/entities"
	"sp/src/usecases/port"

	"gorm.io/gorm"
)

type ContentStorage struct {
}

func NewContentStorage() port.ContentStorage {
	return &ContentStorage{
	}
}

func (pr *ContentStorage) Create(c *entities.Content) (*entities.Content, error) {
	return c, nil
}

func (pr *ContentStorage) Get(id string) (*entities.Content, error) {
	return &entities.Content{
		Content:     []byte{},
		MetaData:    [][]byte{},
		HashedData:  [][]byte{},
		ContentName: "",
		SplitCount:  0,
		Owner:       "",
		Id:          id,
		UserId:      id,
		ContentId:   id,
	}, nil
}

