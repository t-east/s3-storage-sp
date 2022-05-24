package interactor

import (
	"sp/src/core"
	entities "sp/src/domains/entities"
	port "sp/src/usecases/port"
)

type ContentHandler struct {
	Repository      port.ContentRepository
	ContentContract port.ContentContract
	ContentStorage  port.ContentStorage
	UserRepo        port.UserRepository
}

func NewContentInputPort(repository port.ContentRepository, contract port.ContentContract, storage port.ContentStorage, userRepo port.UserRepository) port.ContentInputPort {
	return &ContentHandler{
		Repository:      repository,
		ContentContract: contract,
		ContentStorage:  storage,
		UserRepo:        userRepo,
	}
}

func (c *ContentHandler) Upload(contentInput *entities.Content) (*entities.Receipt, error) {
	//* ブロックチェーンに登録
	err := c.ContentContract.Register(contentInput)
	if err != nil {
		return nil, err
	}
	//* 登録済みユーザかを確認する．
	// _, err = c.UserRepo.FindByID(contentInput.UserId)
	// if err != nil {
	// 	return nil, err
	// }
	//* ulidを作成
	contentInput.Id = core.MakeULID()
	//* コンテンツをストレージに保存
	_, err = c.ContentStorage.Create(contentInput)
	if err != nil {
		return nil, err
	}
	//* データベースに保存
	receipt, err := c.Repository.Create(contentInput)
	if err != nil {
		return nil, err
	}
	return receipt, nil
}

func (c *ContentHandler) FindByID(id string) {
	//* content情報を取得
	_, err := c.Repository.Find(id)
	if err != nil {
		return
	}
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
