package interactor

import (
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
	// TODO データベースからユーザを検索．登録済みアドレスの場合ははじく
	// found, err := uc.Repository.FindByAddress(user.Address)
	// if err != nil {
	// 	uc.OutputPort.RenderError(err)
	// 	return nil, err
	// }
	// if found != nil {
	// 	uc.OutputPort.RenderError(err)
	// 	return nil, err
	// }
	//* データベースに保存
	created, err := uc.Repository.Create(user)
	if err != nil {
		uc.OutputPort.RenderError(err)
		return nil, err
	}
	uc.OutputPort.Render(created, 201)
	return created, nil
}

//* ユーザ情報を取得
func (uc *UserHandler) FindByID(id string) (*entities.User, error) {
	user, err := uc.Repository.FindByID(id)
	if err != nil {
		uc.OutputPort.RenderError(err)
		return nil, err
	}
	uc.OutputPort.Render(user, 200)
	return user, nil
}

// func (uc *UserHandler) KeyGen(id string) (*entities.User, error) {
// 	user, err := uc.Repository.FindByID(id)
// 	if err != nil {
// 		uc.OutputPort.RenderError(err)
// 	}
// 	_, err = uc.Crypt.KeyGen()
// 	if err != nil {
// 		uc.OutputPort.RenderError(err)
// 	}
// 	// TODO; entitiesにkeyを加える処理を付ける
// 	updatedUser := user
// 	updatedUser, err = uc.Repository.Update(updatedUser)
// 	if err != nil {
// 		uc.OutputPort.RenderError(err)
// 		return nil, err
// 	}
// 	uc.OutputPort.Render(updatedUser, 200)
// 	return updatedUser, nil
// }