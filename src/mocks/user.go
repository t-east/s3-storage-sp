package mocks

import (
	entities "sp/src/domains/entities"
	"sp/src/usecases/port"
)

type UserControllerMock struct {
	RepoFactory   func() port.UserRepository
	CryptFactory  func() port.UserCrypt
	OutputFactory func() port.UserOutputPort
	InputFactory  func(
		o port.UserOutputPort,
		u port.UserRepository,
		cr port.UserCrypt,
	) port.UserInputPort
}

func NewUserRepositoryMock() port.UserRepository {
	return &userRepositoryMock{}
}

func NewUserCryptMock() port.UserCrypt {
	return &userCryptMock{}
}

func NewUserOutputPortMock() port.UserOutputPort {
	return &userOutputPortMock{}
}

type userRepositoryMock struct {
}

type userCryptMock struct {
}

type userOutputPortMock struct {
}

func (m *userRepositoryMock) Create(user *entities.User) (*entities.User, error) {
	created := &entities.User{ID: "7", Address: user.Address, PubKey: user.PubKey, PrivKey: user.PrivKey}
	return created, nil
}

func (m *userRepositoryMock) FindByID(id string) (*entities.User, error) {
	user := &entities.User{ID: id, Address: "sdf", PubKey: "sdf", PrivKey: "sdf"}
	return user, nil
}

func (m *userRepositoryMock) Update(user *entities.User) (*entities.User, error) {
	updated := &entities.User{ID: "7", Address: user.Address, PubKey: user.PubKey, PrivKey: user.PrivKey}
	return updated, nil
}

func (m *userCryptMock) KeyGen() (*entities.Key, error) {
	key := &entities.Key{PubKey: "sdf", PrivKey: "sdf"}
	return key, nil
}

func (m *userOutputPortMock) Render(*entities.User, int) {
}

func (m *userOutputPortMock) RenderError(error) {
}
