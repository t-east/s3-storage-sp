package contracts

import (
	"os"
	"sp/src/domains/entities"
	"sp/src/drivers/ethereum"
	"sp/src/usecases/port"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type Param struct {
	Pairing string
	G       []byte
	U       []byte
}

type ContentContract struct{}

func NewContentContracts() port.ContentContract {
	return &ContentContract{}
}

func (cc *ContentContract) Set(content *entities.Content) error {
	privKey := os.Getenv("SP_PRIVATE_KEY")
	conn, client := ethereum.ConnectContentNetWork()
	auth, err := ethereum.AuthUser(client, privKey)
	if err != nil {
		return err
	}
	_, err = conn.SetContentLog(auth, content.HashData, content.ID, common.HexToAddress(content.Address))
	if err != nil {
		return err
	}
	return nil
}

func (cc *ContentContract) Get() ([]*entities.ContentInBlockChain, error) {
	conn, _ := ethereum.ConnectContentNetWork()
	list, err := conn.ListContentLog(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	var logs []*entities.ContentInBlockChain
	for i := 0; i < len(list); i++ {
		logs = append(logs, &entities.ContentInBlockChain{
			HashedData: list[i].Hash,
			ContentId:  list[i].LogId,
		})
	}
	return logs, nil
}
