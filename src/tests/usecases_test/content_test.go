package usecases_test

import (
	"testing"
	"sp/src/domains/entities"
	"sp/src/mocks"
	"sp/src/usecases/interactor"
)

func TestContentUpload(t *testing.T) {
	outputPort := mocks.NewContentOutputPortMock()
	repository := mocks.NewContentRepositoryMock()
	crypt := mocks.NewContentCryptMock()
	sp := mocks.NewContentSPMock()
	contract := mocks.NewContentContractMock()
	inputPort := interactor.NewContentInputPort(outputPort, repository, crypt, sp, contract)

	contentInput := &entities.ContentInput{
		Content:     []byte{},
		ContentName: "コンテンツ1",
		Owner:       "オーナー1",
	}

	receipt, err := inputPort.Upload(contentInput)
	if err != nil {
		t.Fatal(err)
	}
	if receipt.Id != "7" {
		t.Errorf("Content.Upload() should return entities.Content.Id = 7, but got = %s", receipt.Id)
	}
	if receipt.Owner != contentInput.Owner {
		t.Errorf("Content.Upload() should return entities.Content.Owner = %s, but got = %s", contentInput.Owner, receipt.Owner)
	}
	if receipt.ContentName != contentInput.ContentName {
		t.Errorf("Content.Upload() should return entities.Content.ContentName = %s, but got = %s", contentInput.ContentName, receipt.ContentName)
	}
}
