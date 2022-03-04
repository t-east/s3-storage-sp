package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sp/src/asserts"
	"sp/src/domains/entities"
	"sp/src/interfaces/contracts"
	"sp/src/interfaces/controllers"
	"sp/src/interfaces/gateways"
	"testing"

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

// UserName, EmailのあるユーザをPOST -> 201を返すかをテスト
func TestCreateUser(t *testing.T) {
	db, err := LoadTestDB()
	if err != nil {
		t.Errorf("Failed to get DB: %v", err)
		return
	}
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(&entities.User{
		Address: "sdf",
		PubKey:  "pubKey",
		PrivKey: "privKey",
	}); err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodPost, "/api/users", &buf)
	rec := httptest.NewRecorder()
	uc := controllers.LoadUserController(db)
	uc.Dispatch(rec, req)
	asserts.AssertEqual(t, http.StatusCreated, rec.Code, rec.Result().Status)
}

func TestGetUser(t *testing.T) {
	db, err := LoadTestDB()
	if err != nil {
		t.Errorf("Failed to get DB: %v", err)
		return
	}

	ur := gateways.NewUserRepository(db)
	u := &entities.User{
		Address: "sdf",
		PubKey:  "pubKey",
		PrivKey: "privKey",
	}
	user, err := ur.Create(u)
	if err != nil {
		t.Errorf("Can't create user: %v", err)
		return
	}
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/users/%s", user.ID), nil)
	rec := httptest.NewRecorder()
	uc := controllers.LoadUserController(db)
	uc.Dispatch(rec, req)
	asserts.AssertEqual(t, http.StatusOK, rec.Code, rec.Result().Status)
}


func TestGetUserError(t *testing.T) {
	db, err := LoadTestDB()
	if err != nil {
		t.Errorf("Failed to get DB: %v", err)
		return
	}

	ur := gateways.NewUserRepository(db)
	u := &entities.User{
		Address: "sdf",
		PubKey:  "pubKey",
		PrivKey: "privKey",
	}
	_, err = ur.Create(u)
	if err != nil {
		t.Errorf("Can't create user: %v", err)
		return
	}
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/users/%s", "sdsdfsdfsdfsfd"), nil)
	rec := httptest.NewRecorder()
	uc := controllers.LoadUserController(db)
	uc.Dispatch(rec, req)
	asserts.AssertEqual(t, http.StatusBadRequest, rec.Code, rec.Result().Status)
}