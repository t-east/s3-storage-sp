package port

import (
	entities "sp/src/domains/entities"
)

type ContentInputPort interface {
	Upload(content *entities.ContentIn, param *entities.Param) (*entities.Receipt, error)
	FindByID(id string)
}

type ContentRepository interface {
	Create(user *entities.Content) (*entities.Content, error)
	Find(id string) (*entities.Content, error)
	All() ([]*entities.Content, error)
}

type ContentContract interface {
	Set( content *entities.Content ) error
	Get( id string ) (*entities.ContentLog, error )
}

type ContentStorage interface {
	Upload( content *entities.Content) (*entities.Content ,error)
	Get(id string) (*entities.Content ,error)
	GetPreSignedURL(key string) (string, error)
}
