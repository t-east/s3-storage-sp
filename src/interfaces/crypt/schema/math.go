package schema

import (
	"bytes"
	"encoding/gob"
	"sp/src/domains/entities"
)

func SplitContent(c entities.Point, splitCount int) ([][]byte, error) {
	var cb bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&cb) // Will write to network.
	err := enc.Encode(c)
	if err != nil {
		return nil, err
	}
	contentByte := cb.Bytes()
	return SplitSlice(contentByte, splitCount)
}
