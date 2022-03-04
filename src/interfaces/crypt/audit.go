package gateways

import (
	"sp/src/domains/entities"
	"sp/src/interfaces/contracts"
	"sp/src/usecases/port"
)

type auditCrypt struct {
	Param contracts.Param
}

func NewAuditCrypt(param contracts.Param) port.AuditCrypt {
	return &auditCrypt{
		Param: param,
	}
}

func (pr *auditCrypt) AuditProofGen(chal *entities.Chal, content *entities.Content) (*entities.Proof, error) {
	// TODO 実装
	return &entities.Proof{}, nil
}
