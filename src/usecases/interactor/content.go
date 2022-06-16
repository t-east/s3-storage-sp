package interactor

import (
	entities "sp/src/domains/entities"
	port "sp/src/usecases/port"
)

type ContentHandler struct {
	Repository      port.ContentRepository
	ContentContract port.ContentContract
	ContentStorage  port.ContentStorage
}

func NewContentInputPort(repository port.ContentRepository, contract port.ContentContract) port.ContentInputPort {
	return &ContentHandler{
		Repository:      repository,
		ContentContract: contract,
	}
}

func (c *ContentHandler) Upload(ci *entities.ContentIn, param *entities.Param) (*entities.Receipt, error) {
	content := &entities.Content{
		Address:  ci.Address,
		Content:  ci.Content,
		MetaData: ci.MetaData,
	}
	// //* コンテンツからからハッシュ値を生成
	// hash, err := core.HashGen(param, content.Content)
	// if err != nil {
	// 	return nil, err
	// }
	content.HashData = []string{"s", "s", "s"}
	// //* FIWAREに保存
	receipt, err := c.Repository.Create(content)
	if err != nil {
		return nil, err
	}
	// //* ブロックチェーンに登録
	// err = c.ContentContract.Set(content)
	// if err != nil {
	// 	return nil, err
	// }
	// cl, err := c.ContentContract.Get(content.ID)
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

func (c *ContentHandler) FindByID(id string) {
	//* content情報を取得
	_, err := c.Repository.Find(id)
	if err != nil {
		return
	}
}

func (c *ContentHandler) FindAll() ([]*entities.Receipt, error) {
	//* content情報を取得
	receipts, err := c.Repository.All()
	if err != nil {
		return nil, err
	}
	return receipts, nil
}

// TODO keyを使ってファイルのurlを取得する実装
func (c *ContentHandler) GetFileByID(key string) string {
	//* 署名付きurlを返す
	url, err := c.ContentStorage.GetPreSignedURL(key)
	if err != nil {
		return ""
	}
	return url
}
