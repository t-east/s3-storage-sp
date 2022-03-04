package crypt

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"sp/src/domains/entities"
	"sp/src/interfaces/contracts"
	"sp/src/usecases/port"

	"github.com/Nik-U/pbc"
)

type auditCrypt struct {
	Param *contracts.Param
}

func NewAuditCrypt(param *contracts.Param) port.AuditCrypt {
	return &auditCrypt{
		Param: param,
	}
}

func hashChallen(n int, c int, k1, k2 []byte, pairing *pbc.Pairing) ([]int, []*pbc.Element) {

	k1Big := new(big.Int).SetBytes(k1)
	k2Big := new(big.Int).SetBytes(k2)
	nBig := big.NewInt(int64(n))
	var aTable []int
	var vTable []*pbc.Element
	for i := 0; i < n; i++ {
		iBig := big.NewInt(int64(i + 1))
		ik1Big := new(big.Int).Mod(new(big.Int).Mul(iBig, k1Big), nBig)
		aTable = append(aTable, int(ik1Big.Int64()))
	}
	for j := 0; j < c; j++ {
		iBig := big.NewInt(int64(j + 1))
		ik2Big := pairing.NewZr().SetBig(new(big.Int).Mul(iBig, k2Big))
		vTable = append(vTable, ik2Big)
	}
	return aTable, vTable
}

func splitSlice(list []byte, size int) ([][]byte, error) {
	if size <= 0 {
		return nil, fmt.Errorf("size need positive number")
	}
	var result [][]byte
	var tmp = list
	splitNum := len(list) / size
	for i := 0; i < size; i++ {
		if i == size {
			if len(tmp) == 0 {
				break
			}
			r := sha256.Sum256(tmp[:])
			result = append(result, r[:])
		} else {
			r := sha256.Sum256(tmp[0:splitNum])
			result = append(result, r[:])
			tmp = tmp[splitNum:]
		}
	}
	return result, nil
}

func getBinaryBySHA256(s string) []byte {
	r := sha256.Sum256([]byte(s))
	return r[:]
}
func MD5(s string) []byte {
	hash := md5.Sum([]byte(s))
	return []byte(hex.EncodeToString(hash[:]))
}

func (pr *auditCrypt) AuditProofGen(
	chal *entities.Chal,
	content *entities.Content,
	contentLog *entities.Content,
) (*entities.Proof, error) {
	var myu *pbc.Element
	var gamma *pbc.Element
	pairing, _ := pbc.NewPairingFromString(pr.Param.Paring)
	splitedFile, _ := splitSlice(content.Content, content.SplitCount)
	aTable, vTable := hashChallen(contentLog.SplitCount, int(chal.C), chal.K1, chal.K2, pairing)

	var MSum *pbc.Element
	for cIndex := 0; cIndex < int(chal.C); cIndex++ {
		meta := pairing.NewG1().SetBytes(content.MetaData[aTable[cIndex]])
		m := pairing.NewG1().SetFromHash(splitedFile[aTable[cIndex]])
		mm := getBinaryBySHA256(m.X().String())
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
		ContentId: contentLog.ContentId,
	}
	return proof, nil

}
