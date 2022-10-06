package core

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"os"
	"sp/src/domains/entities"

	"github.com/Nik-U/pbc"
)

func MD5(s string) []byte {
	hash := md5.Sum([]byte(s))
	return []byte(hex.EncodeToString(hash[:]))
}

func UseFileRead(fileName string) (*os.File, error) {
	fp, err := os.Open(fileName)
	if err != nil {
		return &os.File{}, err
	}
	return fp, nil
}

func CreateParamMock() (*entities.Param, *entities.Key, error) {
	params := pbc.GenerateA(uint32(160), uint32(512))
	pairing := params.NewPairing()
	g := pairing.NewG1().Rand()
	u := pairing.NewG1().Rand()
	privKey := pairing.NewZr().Rand()
	pubKey := pairing.NewG1().MulZn(g, privKey)
	p := &entities.Param{
		Pairing: params.String(),
		G:       g.Bytes(),
		U:       u.Bytes(),
	}
	k := &entities.Key{
		PubKey:  string(pubKey.Bytes()),
		PrivKey: string(privKey.Bytes()),
	}
	return p, k, nil
}

type A struct {
	Hoge []byte `json:"hoge"`
}

func (a *A) UnmarshalJSON(data []byte) error {
	type A2 struct {
		Hoge string `json:"hoge"`
	}
	a2 := new(A2)
	if err := json.Unmarshal(data, a2); err != nil {
		return err
	}
	a.Hoge = []byte(a2.Hoge)
	return nil
}
