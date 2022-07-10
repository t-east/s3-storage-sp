package crypt

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"math/rand"
	"sp/src/core"
	"sp/src/domains/entities"
	"sp/src/usecases/port"

	"github.com/Nik-U/pbc"
)

type AuditCrypt struct {
	Param *entities.Param
}

func NewAuditCrypt(param *entities.Param) port.AuditCrypt {
	return &AuditCrypt{
		Param: param,
	}
}

func (pr *AuditCrypt) AuditProofGen(
	chal *entities.Chal,
	content *entities.Content,
	contentLog *entities.Content,
) (*entities.Proof, error) {
	var myu *pbc.Element
	var gamma *pbc.Element
	pairing, err := pbc.NewPairingFromString(pr.Param.Pairing)
	if err != nil {
		return nil, err
	}
	var cb bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&cb) // Will write to network.
	err = enc.Encode(content.Content)
	if err != nil {
		log.Fatal("encode error:", err)
	}
	contentByte := cb.Bytes()
	splitedFile, err := core.SplitSlice(contentByte, 3)

	if err != nil {
		return nil, err
	}
	aTable, vTable := core.HashChallen(int(chal.C), []byte(chal.K1), []byte(chal.K2), pairing)
	var MSum *pbc.Element
	if chal.C < 1 {
		return nil, fmt.Errorf("challengeの形が良くない")
	}
	for cIndex := 0; cIndex < int(chal.C); cIndex++ {
		meta := pairing.NewG1().SetBytes([]byte(content.MetaData[aTable[cIndex]]))
		m := pairing.NewG1().SetFromHash(splitedFile[aTable[cIndex]])
		mm := core.GetBinaryBySHA256(m.X().String())
		M := pairing.NewG1().SetBytes(mm)
		if cIndex == 0 {
			myu = pairing.NewZr().MulBig(vTable[cIndex], m.X())
			gamma = pairing.NewG1().PowZn(meta, vTable[cIndex])
		} else {
			myu = pairing.NewZr().Add(myu, pairing.NewZr().MulBig(vTable[cIndex], m.X()))
			gamma.Mul(gamma, pairing.NewG1().PowZn(meta, vTable[cIndex]))
		}
		if cIndex == 0 {
			MSum = pairing.NewG1().PowZn(M, vTable[cIndex])
		} else {
			MSum.Mul(MSum, pairing.NewG1().PowZn(M, vTable[cIndex]))
		}
	}
	proof := &entities.Proof{
		Myu:       myu.Bytes(),
		Gamma:     gamma.Bytes(),
		ContentId: content.ID,
	}
	return proof, nil
}

func AuditChallen(para *entities.Param) *entities.Chal {
	pairing, _ := pbc.NewPairingFromString(para.Pairing)

	ck := rand.Intn(2) + 1
	k1 := pairing.NewZr().Rand()
	k2 := pairing.NewZr().Rand()
	return &entities.Chal{
		ContentId: "",
		C:         ck,
		K1:        k1.Bytes(),
		K2:        k2.Bytes(),
	}
}

func MakeMetaData(uc *entities.ContentInForUser, param *entities.Param) (*entities.Content, error) {
	pairing, err := pbc.NewPairingFromString(param.Pairing)
	if err != nil {
		return nil, err
	}
	var cb bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&cb) // Will write to network.
	err = enc.Encode(uc.Content)
	if err != nil {
		log.Fatal("encode error:", err)
	}
	contentByte := cb.Bytes()
	u := pairing.NewG1().SetBytes(param.U)
	splitCount := 3
	splitedFile, err := core.SplitSlice(contentByte, splitCount)

	if err != nil {
		return nil, err
	}
	privKey := pairing.NewZr().SetBytes([]byte(uc.PrivKey))

	// メタデータの作成
	var metaData [][]byte
	var hash [][]byte
	for i := 0; i < len(splitedFile); i++ {
		m := pairing.NewG1().SetFromHash(splitedFile[i])

		mm := core.GetBinaryBySHA256(m.X().String())
		M := pairing.NewG1().SetBytes(mm)

		um := pairing.NewG1().PowBig(u, m.X())
		temp := pairing.NewG1().Mul(um, M)
		meta := pairing.NewG1().MulZn(temp, privKey)

		metaData = append(metaData, meta.Bytes()[:])
		hash = append(hash, mm)
	}

	return &entities.Content{
		Content:  uc.Content,
		MetaData: metaData,
		HashData: hash,
	}, nil
}

func AuditVerify(params *entities.Param, pubKeyStr string, content *entities.Content, proof *entities.Proof, chal *entities.Chal) ([]byte, []byte, error) {
	pairing, _ := pbc.NewPairingFromString(params.Pairing)
	aTable, vTable := core.HashChallen(chal.C, []byte(chal.K1), []byte(chal.K2), pairing)
	g := pairing.NewG1().SetBytes(params.G)
	pubKey := pairing.NewG1().SetBytes([]byte(pubKeyStr))
	u := pairing.NewG1().SetBytes(params.U)
	myuT := pairing.NewZr().SetBytes([]byte(proof.Myu))
	gammaT := pairing.NewG1().SetBytes([]byte(proof.Gamma))
	var MSum *pbc.Element
	for c := 0; c < chal.C; c++ {
		M := pairing.NewG1().SetBytes([]byte(content.HashData[aTable[c]]))
		if c == 0 {
			MSum = pairing.NewG1().PowZn(M, vTable[c])
		} else {
			MSum.Mul(MSum, pairing.NewG1().PowZn(M, vTable[c]))
		}
	}
	uProof := pairing.NewG1().PowZn(u, myuT)
	right_hand := pairing.NewG1().Mul(uProof, MSum)
	pairingLeft := pairing.NewGT().Pair(gammaT, g)
	pairingRight := pairing.NewGT().Pair(right_hand, pubKey)
	return pairingLeft.Bytes(), pairingRight.Bytes(), nil

}
