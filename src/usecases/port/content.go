package port

import (
	entities "sp/src/domains/entities"
)

type ContentInputPort interface {
	Upload(content *entities.Content) (*entities.Receipt, error)
	FindByID(id string)
}

type ContentOutputPort interface {
	Render(*entities.Receipt, int)
	RenderError(error)
}

type ContentRepository interface {
	Create(user *entities.Content) (*entities.Receipt, error)
	Find(id string) (*entities.Receipt, error)
}

type ContentContract interface {
	Register( content *entities.Content ) error
}
