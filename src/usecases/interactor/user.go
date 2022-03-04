package interactor

import (
	entities "sp/src/domains/entities"
	port "sp/src/usecases/port"
)

type UserHandler struct {
	OutputPort port.UserOutputPort
	Repository port.UserRepository
	Crypt      port.UserCrypt
}

// NewUserInputPort はUserInputPortを取得します．
func NewUserInputPort(outputPort port.UserOutputPort, repository port.UserRepository, crypt port.UserCrypt) port.UserInputPort {
	return &UserHandler{
		OutputPort: outputPort,
		Repository: repository,
		Crypt:      crypt,
	}
}

//* ユーザ登録
func (uc *UserHandler) Create(user *entities.User) (*entities.User, error) {
	//* 鍵生成
	// TODO: ユーザ情報を確認する -> 登録済みなのに新しく鍵を生成したら色々狂う
	key, err := uc.Crypt.KeyGen()
	if err != nil {
		uc.OutputPort.RenderError(err)
		return nil, err
	}
	// TODO: バリデーションを付ける
	user.PubKey = key.PubKey
	user.PrivKey = key.PrivKey

	//* データベースに保存
	user, err = uc.Repository.Create(user)
	if err != nil {
		uc.OutputPort.RenderError(err)
		return nil, err
	}
	uc.OutputPort.Render(user, 201)
	return user, nil
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
