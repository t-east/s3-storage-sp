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

type AuditStorage interface {
	GetContent(string) (*entities.Content, error)
}
