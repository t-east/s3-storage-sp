package port

import (
	entities "sp/src/domains/entities"
)

type AuditInputPort interface {
	ProofGen() (*entities.ProofList, error)
}

type AuditRepository interface {
	Create(proof *entities.Proof) (*entities.Proof, error)
	Update(proof *entities.Proof) (*entities.Proof, error)
}

type AuditContract interface {
	RegisterProof( content *entities.Proof ) error
	GetChallen(string) (*entities.Challenge, error)
	GetAuditLog(id string) (*entities.AuditLog, error)
	Get(string) (*entities.ContentInBlockChain, error)
}

type AuditStorage interface {
	GetContent(string) (*entities.Content, error)
}

type AuditCrypt interface {
	AuditProofGen( chal *entities.Challenge, content *entities.Content, contentLog *entities.ContentInBlockChain ) (*entities.Proof, error)
}