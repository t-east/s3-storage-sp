package gateways

import (
	"sp/src/domains/entities"
	"sp/src/usecases/port"

	"gorm.io/gorm"
)

type ContentSQLHandler interface {
	Find(interface{}, ...interface{}) (*entities.Receipt, error)
	First(interface{}, ...interface{}) (*entities.Receipt, error)
	Create(interface{}) (*entities.Receipt, error)
	Save(interface{}) error
	Delete(interface{}) *entities.Content
	Where(interface{}, ...interface{}) *entities.Content
}

type ContentRepository struct {
	Conn *gorm.DB
	ContentSQLHandler
}

func NewContentRepository(conn *gorm.DB) port.ContentRepository {
	return &ContentRepository{
		Conn: conn,
	}
}

func (ur *ContentRepository) Find(id string) (receipt *entities.Receipt, err error) {
	ContentInDB, err := ur.ContentSQLHandler.Find(&receipt, id)
	if err != nil {
		return nil, err
	}
	return ContentInDB, nil
}

func (ur *ContentRepository) Create(u *entities.Content) (receipt *entities.Receipt, err error) {
	receipt, err = ur.ContentSQLHandler.Create(u)
	if err != nil {
		return nil, err
	}
	return receipt, nil
}
