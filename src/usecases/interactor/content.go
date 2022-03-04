package interactor

import (
	entities "sp/src/domains/entities"
	port "sp/src/usecases/port"
)

type ContentHandler struct {
	OutputPort      port.ContentOutputPort
	Repository      port.ContentRepository
	ContentContract port.ContentContract
}

func NewContentInputPort(outputPort port.ContentOutputPort, repository port.ContentRepository, contract port.ContentContract) port.ContentInputPort {
	return &ContentHandler{
		OutputPort:      outputPort,
		Repository:      repository,
		ContentContract: contract,
	}
}

func (c *ContentHandler) Upload(contentInput *entities.Content) (*entities.Receipt, error) {
	//* ブロックチェーンに登録
	err := c.ContentContract.Register(contentInput)
	if err != nil {
		c.OutputPort.RenderError(err, 500)
		return nil, err
	}

	//* データベースに保存
	receipt, err := c.Repository.Create(contentInput)
	if err != nil {
		c.OutputPort.RenderError(err, 400)
		return nil, err
	}
	c.OutputPort.Render(receipt, 201)
	return receipt, nil
}

func (c *ContentHandler) FindByID(id string) {
	//* content情報を取得
	receipt, err := c.Repository.Find(id)
	if err != nil {
		c.OutputPort.RenderError(err, 400)
		return
	}

	c.OutputPort.Render(receipt, 201)
}
