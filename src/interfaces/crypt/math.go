package crypt

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"math/rand"
	"sp/src/domains/entities"
	"sp/src/interfaces/crypt/schema"
	"sp/src/usecases/port"

	"github.com/Nik-U/pbc"
)

type AuditCrypt struct {
	Param *entities.Param
}

func NewAuditCrypt(param *entities.Param) port.CryptPort {
	return &AuditCrypt{
		Param: param,
	}
}

func (pr *AuditCrypt) ContentHashGen(c *entities.Content) (*entities.Content, error) {
	param, err := schema.GeneratePairingParam(pr.Param)
	if err != nil {
		return nil, err
	}

	splitCount := 3
	splitFile, err := schema.SplitContent(c.Content, splitCount)
	if err != nil {
		return nil, err
	}

	var hash [][]byte
	for i := 0; i < len(splitFile); i++ {
		m := param.Pairing.NewG1().SetFromHash(splitFile[i])

		mm := schema.GetBinaryBySHA256(m.X().String())
		hash = append(hash, mm)
	}

	c.HashData = hash
	return c, nil
}

func (pr *AuditCrypt) AuditProofGen(
	challenge *entities.Challenge,
	content *entities.Content,
	contentLog *entities.ContentInBlockChain,
) (*entities.Proof, error) {
	param, err := schema.GeneratePairingParam(pr.Param)
	if err != nil {
		return nil, err
	}

	splitCount := 3
	splitFile, err := schema.SplitContent(content.Content, splitCount)
	if err != nil {
		return nil, err
	}

	var myu *pbc.Element
	var gamma *pbc.Element

	if err != nil {
		return nil, err
	}
	aTable, vTable := schema.HashChallen(int(challenge.C), []byte(challenge.K1), []byte(challenge.K2), param.Pairing)
	var MSum *pbc.Element
	if challenge.C < 1 {
		return nil, fmt.Errorf("challengeの形が良くない")
	}
	for cIndex := 0; cIndex < int(challenge.C); cIndex++ {
		meta := param.Pairing.NewG1().SetBytes([]byte(content.MetaData[aTable[cIndex]]))
		m := param.Pairing.NewG1().SetFromHash(splitFile[aTable[cIndex]])
		mm := schema.GetBinaryBySHA256(m.X().String())
		M := param.Pairing.NewG1().SetBytes(mm)
		if cIndex == 0 {
			myu = param.Pairing.NewZr().MulBig(vTable[cIndex], m.X())
			gamma = param.Pairing.NewG1().PowZn(meta, vTable[cIndex])
		} else {
			myu = param.Pairing.NewZr().Add(myu, param.Pairing.NewZr().MulBig(vTable[cIndex], m.X()))
			gamma.Mul(gamma, param.Pairing.NewG1().PowZn(meta, vTable[cIndex]))
		}
		if cIndex == 0 {
			MSum = param.Pairing.NewG1().PowZn(M, vTable[cIndex])
		} else {
			MSum.Mul(MSum, param.Pairing.NewG1().PowZn(M, vTable[cIndex]))
		}
	}
	proof := &entities.Proof{
		Myu:       myu.Bytes(),
		Gamma:     gamma.Bytes(),
		ContentId: content.ID,
	}
	return proof, nil
}

func AuditChallen(para *entities.Param) (*entities.Challenge, error) {
	param, err := schema.GeneratePairingParam(para)
	if err != nil {
		return nil, err
	}

	ck := rand.Intn(2) + 1
	k1 := param.Pairing.NewZr().Rand()
	k2 := param.Pairing.NewZr().Rand()
	return &entities.Challenge{
		ContentId: "",
		C:         ck,
		K1:        k1.Bytes(),
		K2:        k2.Bytes(),
	}, nil
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
	splitedFile, err := schema.SplitSlice(contentByte, splitCount)

	if err != nil {
		return nil, err
	}
	privKey := pairing.NewZr().SetBytes([]byte(uc.PrivKey))

	// メタデータの作成
	var metaData [][]byte
	var hash [][]byte
	for i := 0; i < len(splitedFile); i++ {
		m := pairing.NewG1().SetFromHash(splitedFile[i])

		mm := schema.GetBinaryBySHA256(m.X().String())
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

func AuditVerify(params *entities.Param, pubKeyStr string, content *entities.Content, proof *entities.Proof, challenge *entities.Challenge) ([]byte, []byte, error) {
	pairing, _ := pbc.NewPairingFromString(params.Pairing)
	aTable, vTable := schema.HashChallen(challenge.C, []byte(challenge.K1), []byte(challenge.K2), pairing)
	g := pairing.NewG1().SetBytes(params.G)
	pubKey := pairing.NewG1().SetBytes([]byte(pubKeyStr))
	u := pairing.NewG1().SetBytes(params.U)
	myuT := pairing.NewZr().SetBytes([]byte(proof.Myu))
	gammaT := pairing.NewG1().SetBytes([]byte(proof.Gamma))
	var MSum *pbc.Element
	for c := 0; c < challenge.C; c++ {
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
