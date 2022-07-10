package interactor

import (
	"errors"
	entities "sp/src/domains/entities"
	port "sp/src/usecases/port"
)

type ContentUseCase struct {
	ContentContract port.ContentContract
	ContentRepo     port.ContentRepository
	ContentCrypt    port.ContentCrypt
}

func NewContentUseCase(contentContract port.ContentContract, contentRepo port.ContentRepository, contenCrypt port.ContentCrypt) *ContentUseCase {
	return &ContentUseCase{
		ContentContract: contentContract,
		ContentRepo:     contentRepo,
		ContentCrypt:    contenCrypt,
	}
}

func (c *ContentUseCase) Upload(ci *entities.ContentIn) (*entities.Receipt, error) {
	content := &entities.Content{
		Address:  ci.Address,
		Content:  ci.Content,
		MetaData: ci.MetaData,
	}
	//* コンテンツからからハッシュ値を生成
	content, err := c.ContentCrypt.ContentHashGen(content)
	if err != nil {
		return nil, err
	}
	// //* FIWAREに保存
	receipt, err := c.ContentRepo.Create(content)
	if err != nil {
		return nil, errors.New("fiware error")
	}
	//* ブロックチェーンに登録
	err = c.ContentContract.Set(content)
	if err != nil {
		return nil, err
	}
	result := &entities.Receipt{
		ID:       receipt.ID,
		Content:  receipt.Content,
		MetaData: receipt.MetaData,
		HashData: receipt.HashData,
	}
	return result, nil
}

func (c *ContentUseCase) FindByID(id string) {
	//* content情報を取得
	_, err := c.ContentRepo.Find(id)
	if err != nil {
		return
	}
}

func (c *ContentUseCase) FindAll() ([]*entities.Receipt, error) {
	//* content情報を取得
	receipts, err := c.ContentRepo.All()
	if err != nil {
		return nil, err
	}
	return receipts, nil
}
