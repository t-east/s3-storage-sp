package crypt

import (
	"bytes"
	"encoding/gob"
	"log"
	"sp/src/core"
	"sp/src/domains/entities"
	"sp/src/usecases/port"

	"github.com/Nik-U/pbc"
)

type ContentCrypt struct {
	Param *entities.Param
}

func NewContentCrypt(param *entities.Param) port.ContentCrypt {
	return &ContentCrypt{
		Param: param,
	}
}

func (cc *ContentCrypt) ContentHashGen(c *entities.Content) (*entities.Content, error) {
	pairing, err := pbc.NewPairingFromString(cc.Param.Pairing)
	if err != nil {
		return nil, err
	}
	var cb bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&cb) // Will write to network.
	err = enc.Encode(c.Content)
	if err != nil {
		log.Fatal("encode error:", err)
	}
	contentByte := cb.Bytes()
	splitCount := 3
	splitedFile, err := core.SplitSlice(contentByte, splitCount)

	if err != nil {
		return nil, err
	}

	var hash []string
	for i := 0; i < len(splitedFile); i++ {
		m := pairing.NewG1().SetFromHash(splitedFile[i])

		mm := core.GetBinaryBySHA256(m.X().String())
		hash = append(hash, string(mm))
	}

	c.HashData = hash
	return c, nil
}
