package interactor

import (
	"errors"
	entities "sp/src/domains/entities"
	port "sp/src/usecases/port"
)

type ContentUseCase struct {
	ContractPort port.ContentContractPort
	ContentRepo  port.ContentRepository
	cryptPort    port.CryptPort
}

func NewContentUseCase(ContractPort port.ContentContractPort, contentRepo port.ContentRepository, contenCrypt port.CryptPort) *ContentUseCase {
	return &ContentUseCase{
		ContractPort: ContractPort,
		ContentRepo:  contentRepo,
		cryptPort:    contenCrypt,
	}
}

func (c *ContentUseCase) Upload(ci *entities.ContentIn) (*entities.Content, error) {
	content := &entities.Content{
		Address:  ci.Address,
		Content:  ci.Content,
		MetaData: ci.MetaData,
	}
	//* コンテンツからからハッシュ値を生成
	content, err := c.cryptPort.ContentHashGen(content)
	if err != nil {
		return nil, err
	}
	// //* FIWAREに保存
	created, err := c.ContentRepo.Create(content)
	if err != nil {
		return nil, errors.New("fiware error")
	}
	//* ブロックチェーンに登録
	err = c.ContractPort.Set(content)
	if err != nil {
		return nil, err
	}
	result := &entities.Content{
		ID:       created.ID,
		Content:  created.Content,
		MetaData: created.MetaData,
		HashData: created.HashData,
	}
	return result, nil
}

func (c *ContentUseCase) FindByID(id string) {
	//* content情報を取得
	_, err := c.ContentRepo.FindByID(id)
	if err != nil {
		return
	}
}

func (c *ContentUseCase) FindAll() ([]*entities.Content, error) {
	//* content情報を取得
	list, err := c.ContentRepo.List()
	if err != nil {
		return nil, err
	}
	return list, nil
}
