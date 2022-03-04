package interactor

import (
	"sp/src/core"
	entities "sp/src/domains/entities"
	port "sp/src/usecases/port"
)

type UserHandler struct {
	OutputPort port.UserOutputPort
	Repository port.UserRepository
}

// NewUserInputPort はUserInputPortを取得します．
func NewUserInputPort(outputPort port.UserOutputPort, repository port.UserRepository) port.UserInputPort {
	return &UserHandler{
		OutputPort: outputPort,
		Repository: repository,
	}
}

//* ユーザ登録
func (uc *UserHandler) Create(user *entities.User) (*entities.User, error) {
	//* データベースからユーザを検索．登録済みアドレスの場合ははじく
	found, err := uc.Repository.FindByAddress(user.Address)
	if err != nil {
		uc.OutputPort.RenderError(err, 400)
		return nil, err
	}
	if found != nil {
		uc.OutputPort.RenderError(err, 400)
		return nil, err
	}
	user.ID = core.MakeULID()
	//* データベースに保存
	created, err := uc.Repository.Create(user)
	if err != nil {
		uc.OutputPort.RenderError(err, 400)
		return nil, err
	}
	uc.OutputPort.Render(created, 201)
	return created, nil
}

//* ユーザ情報を取得
func (uc *UserHandler) FindByID(id string) (*entities.User, error) {
	user, err := uc.Repository.FindByID(id)
	if err != nil {
		uc.OutputPort.RenderError(err, 400)
		return nil, err
	}
	uc.OutputPort.Render(user, 200)
	return user, nil
}