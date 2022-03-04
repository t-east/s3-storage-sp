package mocks

import (
	entities "sp/src/domains/entities"
	"sp/src/usecases/port"
)

type ContentControllerMock struct {
	RepoFactory   func() port.ContentRepository
	CryptFactory  func() port.ContentCrypt
	OutputFactory func() port.ContentOutputPort
	InputFactory  func(
		o port.ContentOutputPort,
		u port.ContentRepository,
		cr port.ContentCrypt,
	) port.ContentInputPort
}

func NewContentRepositoryMock() port.ContentRepository {
	return &ContentRepositoryMock{}
}

func NewContentCryptMock() port.ContentCrypt {
	return &ContentCryptMock{}
}

func NewContentOutputPortMock() port.ContentOutputPort {
	return &ContentOutputPortMock{}
}

func NewContentSPMock() port.ContentSP {
	return &ContentSPMock{}
}

func NewContentContractMock() port.ContentContract {
	return &ContentContractMock{}
}

type ContentRepositoryMock struct {
}

type ContentCryptMock struct {
}

type ContentOutputPortMock struct {
}

type ContentSPMock struct {
}

type ContentContractMock struct {
}

func (m *ContentRepositoryMock) Create(content *entities.Content) (*entities.ContentInDB, error) {
	created := &entities.ContentInDB{
		Id:          "7",
		UserId:      content.UserId,
		ContentName: content.ContentName,
		Owner:       content.Owner,
	}
	return created, nil
}

func (m *ContentRepositoryMock) Find(id string) (*entities.Content, error) {
	found := &entities.Content{Id: id, UserId: "1"}
	return found, nil
}

func (m *ContentCryptMock) MakeMetaData(input *entities.ContentInput) (*entities.Content, error) {
	content := &entities.Content{
		Content:     input.Content,
		MetaData:    [][]byte{},
		HashedData:  [][]byte{},
		ContentName: input.ContentName,
		SplitCount:  0,
		Owner:       input.Owner,
		Id:          "",
		UserId:      "",
	}
	return content, nil
}

func (m *ContentOutputPortMock) Render(*entities.ReceiptFromSP, int) {
}

func (m *ContentOutputPortMock) RenderError(error) {
}

func (m *ContentSPMock) UploadSP(content *entities.Content) (*entities.ReceiptFromSP, error) {
	receipt := &entities.ReceiptFromSP{
		ContentName: content.ContentName,
		Owner:       content.Owner,
		Id:          content.Id,
	}
	return receipt, nil
}

func (m *ContentSPMock) GetContent(id string) (*entities.ReceiptFromSP, error) {
	receipt := &entities.ReceiptFromSP{
		ContentName: "コンテンツ1",
		Owner:       "オーナー",
		Id:          id,
	}
	return receipt, nil
}

func (m *ContentContractMock) Register(string, string) error {
	return nil
}
