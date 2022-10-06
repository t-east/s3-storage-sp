package schema

import (
	"sp/src/domains/entities"

	"github.com/Nik-U/pbc"
)

type ParingParam struct {
	Pairing *pbc.Pairing
	U       *pbc.Element
	G       *pbc.Element
}

type PairingKey struct {
	PrivKey *pbc.Element
	PubKey  *pbc.Element
}

func GeneratePairingParam(p *entities.Param) (*ParingParam, error) {
	pairing, err := pbc.NewPairingFromString(p.Pairing)
	if err != nil {
		return nil, err
	}
	u := pairing.NewG1().SetBytes([]byte(p.U))
	g := pairing.NewG1().SetBytes([]byte(p.G))
	return &ParingParam{
		Pairing: pairing,
		U:       u,
		G:       g,
	}, nil
}

func (p *ParingParam) GeneratePairingKey(prk []byte) (*PairingKey, error) {
	privKey := p.Pairing.NewZr().SetBytes(prk)
	pubKey := p.Pairing.NewG1().MulZn(p.G, privKey)
	return &PairingKey{
		PrivKey: privKey,
		PubKey:  pubKey,
	}, nil
}
