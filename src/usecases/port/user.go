package port

import (
	entities "sp/src/domains/entities"
)

type UserInputPort interface {
	Create(*entities.User) (*entities.User, error)
	FindByID(uint) (*entities.User, error)
}

type UserRepository interface {
	Create(*entities.User) (*entities.User, error)
	Update(*entities.User) (*entities.User, error)
	FindByID(uint) (*entities.User, error)
	FindByAddress(string) (*entities.User, error)
}
