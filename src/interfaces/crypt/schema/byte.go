package schema

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"math/big"
	"sp/src/domains/entities"

	"github.com/Nik-U/pbc"
)

func HashGen(param *entities.Param, content entities.Point) ([]string, error) {
	pairing, err := pbc.NewPairingFromString(param.Pairing)
	if err != nil {
		return nil, err
	}
	splitCount := 3
	var cb bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&cb) // Will write to network.
	err = enc.Encode(content)
	if err != nil {
		return nil, err
	}
	contentByte := cb.Bytes()
	splitedFile, err := SplitSlice(contentByte, splitCount)
	if err != nil {
		return nil, err
	}
	var hashData []string
	for i := 0; i < len(splitedFile); i++ {
		m := pairing.NewG1().SetFromHash(splitedFile[i])
		mm := m.X().String()
		hashData = append(hashData, mm)
	}
	return hashData, nil
}

func HashChallen(c int, k1, k2 []byte, pairing *pbc.Pairing) ([]int, []*pbc.Element) {
	n := 3
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

func SplitSlice(list []byte, size int) ([][]byte, error) {
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

func GetBinaryBySHA256(s string) []byte {
	r := sha256.Sum256([]byte(s))
	return r[:]
}
