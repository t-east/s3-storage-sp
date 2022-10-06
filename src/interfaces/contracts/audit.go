package contracts

import (
	"os"
	"sp/src/domains/entities"
	"sp/src/drivers/ethereum"
	"sp/src/usecases/port"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type AuditContract struct{}

func NewAuditContracts() port.AuditContract {
	return &AuditContract{}
}

func (cc *AuditContract) GetChallen(id string) (*entities.Challenge, error) {

	conn, _ := ethereum.ConnectAuditNetWork()
	log, err := conn.GetAuditLog(&bind.CallOpts{}, id)
	if err != nil {
		return nil, err
	}
	return &entities.Challenge{
		ContentId: id,
		C:         int(log.Chal),
		K1:        log.K1,
		K2:        log.K2,
	}, nil
}

func (cc *AuditContract) RegisterProof(proof *entities.Proof) error {
	privKey := os.Getenv("SP_PRIVATE_KEY")
	conn, client := ethereum.ConnectAuditNetWork()
	auth, err := ethereum.AuthUser(client, privKey)
	if err != nil {
		return err
	}
	_, err = conn.SetProof(auth, proof.ContentId, proof.Myu, proof.Gamma)
	return err
}

func (cc *AuditContract) GetAuditLog(id string) (*entities.AuditLog, error) {
	conn, _ := ethereum.ConnectAuditNetWork()
	a, err := conn.GetAuditLog(&bind.CallOpts{}, id)
	if err != nil {
		return nil, err
	}
	return &entities.AuditLog{
		Chal: &entities.Challenge{
			ContentId: id,
			C:         int(a.Chal),
			K1:        a.K1,
			K2:        a.K2,
		},
		Proof: &entities.Proof{
			Myu:       a.Myu,
			Gamma:     a.Gamma,
			ContentId: id,
		},
		Result:    a.Result,
		ContentID: id,
	}, nil
}

func (cc *AuditContract) Get(id string) (*entities.ContentInBlockChain, error) {
	conn, _ := ethereum.ConnectContentNetWork()
	l, err := conn.GetContentLog(&bind.CallOpts{}, id)
	if err != nil {
		return nil, err
	}
	return &entities.ContentInBlockChain{
		HashedData: l.Hash,
		ContentId:  id,
	}, nil
}
