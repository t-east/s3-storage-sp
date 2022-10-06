package port

import (
	entities "sp/src/domains/entities"
)

type CryptPort interface {
	ContentHashGen(content *entities.Content) (*entities.Content, error)
	AuditProofGen(chal *entities.Challenge, content *entities.Content, contentLog *entities.ContentInBlockChain) (*entities.Proof, error)
}
