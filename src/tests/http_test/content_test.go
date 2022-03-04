package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sp/src/asserts"
	"sp/src/domains/entities"
	"sp/src/interfaces/controllers"
	"sp/src/interfaces/gateways"
	"testing"
)

//* 登録済みユーザのuserIdでコンテンツを作成
func TestCreateContent(t *testing.T) {
	db, err := LoadTestDB()
	if err != nil {
		t.Errorf("Failed to get DB: %v", err)
		return
	}
	ur := gateways.NewUserRepository(db)
	u := &entities.User{
		Address: "sdf",
		PubKey:  []byte("pubKey"),
		PrivKey: []byte("privKey"),
	}
	user, _ := ur.Create(u)
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(&entities.Content{
		Content:     []byte{1,2,3,4,5},
		MetaData:    [][]byte{},
		HashedData:  [][]byte{},
		ContentName: "コンテンツ1",
		SplitCount:  2,
		Owner:       "sdf",
		Id:          "sdf",
		UserId:      user.ID,
		ContentId:   "sdf",
	}); err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodPost, "/api/content", &buf)
	rec := httptest.NewRecorder()
	cc := controllers.LoadContentController(db)
	cc.Dispatch(rec, req)
	asserts.AssertEqual(t, http.StatusCreated, rec.Code, rec.Result().Status)
}

//* 登録済みユーザのuserIdでコンテンツを取得
func TestGetContent(t *testing.T) {
	db, err := LoadTestDB()
	if err != nil {
		t.Errorf("Failed to get DB: %v", err)
		return
	}

	ur := gateways.NewUserRepository(db)
	u := &entities.User{
		Address: "sdf",
		PubKey: []byte("pubKey"),
		PrivKey: []byte("privKey"),
	}
	user, err := ur.Create(u)
	if err != nil {
		t.Errorf("Can't create user: %v", err)
		return
	}
	cr := gateways.NewContentRepository(db)
	c := &entities.Content{
		Content:     []byte{},
		MetaData:    [][]byte{},
		HashedData:  [][]byte{},
		ContentName: "",
		SplitCount:  0,
		Owner:       "",
		Id:          "sdf",
		UserId:      user.ID,
	}
	content, err := cr.Create(c)
	if err != nil {
		t.Errorf("Can't create user: %v", err)
		return
	}
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/contents/%s", content.Id), nil)
	rec := httptest.NewRecorder()
	cc := controllers.LoadContentController(db)
	cc.Dispatch(rec, req)
	asserts.AssertEqual(t, http.StatusOK, rec.Code, rec.Result().Status)
}

//* 存在しないUserIDで作成する -> エラー
func TestGetContentError(t *testing.T) {
	db, err := LoadTestDB()
	if err != nil {
		t.Errorf("Failed to get DB: %v", err)
		return
	}
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/contents/%s", "sdfsdf"), nil)
	rec := httptest.NewRecorder()
	uc := controllers.LoadUserController(db)
	uc.Dispatch(rec, req)
	asserts.AssertEqual(t, http.StatusBadRequest, rec.Code, rec.Result().Status)
}
