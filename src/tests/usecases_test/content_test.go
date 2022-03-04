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
	contract := mocks.NewContentContractMock()
	inputPort := interactor.NewContentInputPort(outputPort, repository, contract)

	contentInput := &entities.Content{
		Content:     []byte{},
		MetaData:    [][]byte{},
		HashedData:  [][]byte{},
		ContentName: "コンテンツ1",
		SplitCount:  0,
		Owner:       "オーナー1",
		Id:          "12",
		UserId:      "1",
	}

	receipt, err := inputPort.Upload(contentInput)
	if err != nil {
		t.Fatal(err)
	}
	if receipt.Id != contentInput.Id {
		t.Errorf("Content.Upload() should return entities.Receipt.Id = 7, but got = %s", receipt.Id)
	}
}