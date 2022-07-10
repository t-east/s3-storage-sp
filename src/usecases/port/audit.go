package port

import (
	entities "sp/src/domains/entities"
)

type AuditInputPort interface {
	ProofGen() (*entities.Proofs, error)
}

type AuditRepository interface {
	Create(proof *entities.Proof) (*entities.Proof, error)
	Update(proof *entities.Proof) (*entities.Proof, error)
}

type AuditContract interface {
	RegisterProof( content *entities.Proof ) error
	GetChallen(string) (*entities.Chal, error)
	GetAuditLog(id string) (*entities.AuditLog, error)
	Get(string) (*entities.ContentInBlockChain, error)
}

type AuditStorage interface {
	GetContent(string) (*entities.Content, error)
}

type AuditCrypt interface {
	AuditProofGen( chal *entities.Chal, content *entities.Receipt, contentLog *entities.ContentInBlockChain ) (*entities.Proof, error)
}