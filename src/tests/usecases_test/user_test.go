package usecases_test

import (
	"testing"
	"sp/src/domains/entities"
	"sp/src/mocks"
	"sp/src/usecases/interactor"
)

func TestUserCreate(t *testing.T) {
	outputPort := mocks.NewUserOutputPortMock()
	repository := mocks.NewUserRepositoryMock()
	crypt := mocks.NewUserCryptMock()
	inputPort := interactor.NewUserInputPort(outputPort, repository, crypt)

	user := &entities.User{Address: "sdf", PubKey: "sdf", PrivKey: "sdf"}

	created, err := inputPort.Create(user)
	if err != nil {
		t.Fatal(err)
	}
	if created.ID != "7" {
		t.Errorf("User.Create() should return entities.User.ID = 7, but got = %s", created.ID)
	}
	if created.Address != user.Address {
		t.Errorf("User.Create() should return entities.User.Address = %s, but got = %s", user.Address, created.Address)
	}
	if created.PubKey != user.PubKey {
		t.Errorf("User.Create() should return entities.User.PubKey = %s, but got = %s", user.PubKey, created.PubKey)
	}
	if created.PrivKey != user.PrivKey {
		t.Errorf("User.Create() should return entities.User.PrivKey = %s, but got = %s", user.PrivKey, created.PrivKey)
	}
}

func TestUserFindByID(t *testing.T) {
	outputPort := mocks.NewUserOutputPortMock()
	repository := mocks.NewUserRepositoryMock()
	crypt := mocks.NewUserCryptMock()
	inputPort := interactor.NewUserInputPort(outputPort, repository, crypt)

	user := &entities.User{Address: "sdf", PubKey: "sdf", PrivKey: "sdf"}

	id := "7"
	found, err := inputPort.FindByID(id)
	if err != nil {
		t.Fatal(err)
	}
	if found.ID != "7" {
		t.Errorf("User.Create() should return entities.User.ID = 7, but got = %s", found.ID)
	}
	if found.Address != user.Address {
		t.Errorf("User.Create() should return entities.User.Address = %s, but got = %s", user.Address, found.Address)
	}
	if found.PubKey != user.PubKey {
		t.Errorf("User.Create() should return entities.User.PubKey = %s, but got = %s", user.PubKey, found.PubKey)
	}
	if found.PrivKey != user.PrivKey {
		t.Errorf("User.Create() should return entities.User.PrivKey = %s, but got = %s", user.PrivKey, found.PrivKey)
	}
}
