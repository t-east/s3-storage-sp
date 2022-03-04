package usecases_test

import (
	"sp/src/domains/entities"
	"sp/src/mocks"
	"testing"
)

// コンテンツ情報永続化モックのテスト
func TestContentRepositoryCreate(t *testing.T) {
	FakeRepo := mocks.NewContentRepositoryMock()
	contentInput := &entities.Content{
		Content:     []byte{},
		MetaData:    [][]byte{},
		HashedData:  [][]byte{},
		ContentName: "コンテンツ1",
		SplitCount:  0,
		Owner:       "オーナー",
		Id:          "",
		UserId:      "12",
	}
	content, err := FakeRepo.Create(contentInput)
	if err != nil {
		t.Fatal(err)
	}
	if content.Id != contentInput.Id {
		t.Errorf("Content.Create() should return entities.Content.Id = %s, but got = %s", contentInput.Id, content.Id)
	}
}

// ブロックチェーン登録モックのテスト
func TestContentContractRegister(t *testing.T) {
	FakeContract := mocks.NewContentContractMock()
	content := &entities.Content{
		Content:     []byte{},
		MetaData:    [][]byte{},
		HashedData:  [][]byte{},
		ContentName: "",
		SplitCount:  0,
		Owner:       "",
		Id:          "",
		UserId:      "",
	}
	err := FakeContract.Register(content)
	if err != nil {
		t.Fatal(err)
	}
}
