package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sp/src/asserts"
	"sp/src/domains/entities"
	"sp/src/interfaces/contracts"
	"sp/src/interfaces/controllers"
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
		ID:      "1",
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
