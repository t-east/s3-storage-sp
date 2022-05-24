package interactor

import (
	"log"
	"sp/src/core"
	entities "sp/src/domains/entities"
	port "sp/src/usecases/port"
	"strconv"
)

type UserHandler struct {
	Repository port.UserRepository
}

// NewUserInputPort はUserInputPortを取得します．
func NewUserInputPort(repository port.UserRepository) port.UserInputPort {
	return &UserHandler{
		Repository: repository,
	}
}

//* ユーザ登録
func (uc *UserHandler) Create(user *entities.User) (*entities.User, error) {
	//* データベースからユーザを検索．登録済みアドレスの場合ははじく
	// TODO 登録済みの場合の処理
	found, _ := uc.Repository.FindByAddress(user.Address)
	// if found != nil {
	// 	uc.OutputPort.RenderError(err, 400)
	// 	return nil, err
	// }
	log.Print(found)
	id64, err := strconv.ParseUint(core.MakeULID(),10, 64)
	if err != nil {
		return nil, err
	}
	user.ID = uint(id64)
	//* データベースに保存
	created, err := uc.Repository.Create(user)
	if err != nil {
		return nil, err
	}
	return created, nil
}

//* ユーザ情報を取得
func (uc *UserHandler) FindByID(id uint) (*entities.User, error) {
	user, err := uc.Repository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
