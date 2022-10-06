package port

import entities "sp/src/domains/entities"

type AuditContractPort interface {
	RegisterProof(content *entities.Proof) error
	GetChallen(string) (*entities.Challenge, error)
	FindByID(auditID string) (*entities.AuditLog, error)
}

type ContentContractPort interface {
	Set(content *entities.Content) error
	List() ([]*entities.ContentInBlockChain, error)
	FindByID(contentID string) (*entities.ContentInBlockChain, error)
}
