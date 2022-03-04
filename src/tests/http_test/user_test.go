package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"sp/src/domains/entities"
	"sp/src/interfaces/contracts"
	"sp/src/interfaces/controllers"
	"sp/src/mocks"
	"sp/src/usecases/interactor"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupUserTestController() (*controllers.UserController, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&entities.User{})
	uc := &controllers.UserController{
		OutputFactory: mocks.NewUserOutputPortMock(),
		InputFactory:  interactorwddddddddddddzeeeeccc.NewUserInputPor,
		RepoFactory:   mocks.NewUserRepository,
		Param:         contracts.Param,
		Conn:          db,
	}
	return uc, nil
}

// UserName, EmailのあるユーザをPOST -> 201を返すかをテスト
func TestCreateUser(t *testing.T) {
	uc, err := SetupUserTestController()
	if err != nil {
		t.Errorf("Failed to get DB: %v", err)
		return
	}
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(&handler.RequestUser{
		UserName: "test",
		Email:    "test@test.com",
	}); err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodPost, "/api/users", &buf)
	rec := httptest.NewRecorder()
	uc.Post(rec, req)
	asserts.AssertEqual(t, http.StatusCreated, rec.Code, rec.Result().Status)
}

// UserNameが無いユーザをPOST -> 400を返すかをテスト
func TestCreateUserError(t *testing.T) {
	uc, err := SetupUserTestController()
	if err != nil {
		t.Errorf("Failed to get DB: %v", err)
		return
	}
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(&handler.RequestUser{
		UserName: "",
		Email:    "test@test.com",
	}); err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodPost, "/api/users/", &buf)
	rec := httptest.NewRecorder()
	uc.Post(rec, req)
	asserts.AssertEqual(t, http.StatusBadRequest, rec.Code, rec.Result().Status)
}
