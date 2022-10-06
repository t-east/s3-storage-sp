package port

import (
	entities "sp/src/domains/entities"
)

type ContentRepository interface {
	Create(user *entities.Content) (*entities.Content, error)
	FindByID(id string) (*entities.Content, error)
	List() ([]*entities.Content, error)
}

type ContentStorage interface {
	Upload(content *entities.Content) (*entities.Content, error)
	FindByID(id string) (*entities.Content, error)
	GetPreSignedURL(key string) (string, error)
}
