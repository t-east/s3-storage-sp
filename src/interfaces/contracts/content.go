package contracts

import (
	"fmt"
	"os"
	"sp/src/domains/entities"
	"sp/src/drivers/ethereum"
	"sp/src/usecases/port"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"
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
	err := godotenv.Load(fmt.Sprintf("./%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		return err
	}
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

func (cc *ContentContract) Get(id string) (*entities.ContentLog, error) {
	err := godotenv.Load(fmt.Sprintf("./%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		return nil, err
	}
	conn, _ := ethereum.ConnectContentNetWork()
	log, err := conn.GetContentLog(&bind.CallOpts{}, id)
	if err != nil {
		return nil, err
	}
	contentLog := &entities.ContentLog{
		Owner:    log.Owner.String(),
		Hash:     log.Hash,
		Provider: log.Provider.String(),
	}
	return contentLog, nil
}