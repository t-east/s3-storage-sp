package usecases_test

import (
	"testing"
	"sp/src/domains/entities"
	"sp/src/mocks"
)

// メタデータ作成モックのテスト
func TestContentMakeMetaData(t *testing.T) {
	FakeCrypt := mocks.NewContentCryptMock()
	contentInput := &entities.ContentInput{
		Content:     []byte{},
		ContentName: "コンテンツ1",
		Owner:       "オーナー1",
	}
	content, err := FakeCrypt.MakeMetaData(contentInput)
	if err != nil {
		t.Fatal(err)
	}
	if content.ContentName != "contentInput.ContentName" {
		t.Errorf("Content.MakeMetaData() should return entities.Content.ContentName = %s, but got = %s", contentInput.ContentName, content.ContentName)
	}
}

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
		UserId:      "12",
	}
	content, err := FakeRepo.Create(contentInput)
	if err != nil {
		t.Fatal(err)
	}
	if content.ContentName != contentInput.ContentName {
		t.Errorf("Content.Create() should return entities.Content.ContentName = %s, but got = %s", contentInput.ContentName, content.ContentName)
	}
	if content.Owner != contentInput.Owner {
		t.Errorf("Content.Create() should return entities.Content.Owner = %s, but got = %s", contentInput.Owner, content.Owner)
	}
	if content.UserId != contentInput.UserId {
		t.Errorf("Content.Create() should return entities.Content.UserId = %s, but got = %s", contentInput.UserId, content.UserId)
	}
	if content.Id != "7" {
		t.Errorf("Content.Create() should return entities.Content.Id = %s, but got = %s", contentInput.Id, content.Id)
	}
}


// ブロックチェーン登録モックのテスト
func TestContentContractRegister(t *testing.T) {
	FakeContract := mocks.NewContentContractMock()
	a := "a"
	b := "b"
	err := FakeContract.Register(a,b)
	if err != nil {
		t.Fatal(err)
	}
}