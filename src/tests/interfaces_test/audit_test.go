package interfaces_test

import (
	"sp/src/domains/entities"
	"sp/src/interfaces/contracts"
	"sp/src/interfaces/crypt"
	"testing"
)

// AuditProofGenテスト
func TestAuditProofGen(t *testing.T) {
	param := &contracts.Param{
		Paring: "",
		G:      []byte{},
		U:      []byte{},
	}
	Pg := crypt.NewAuditCrypt(param)
	testByte := []byte{1}
	contentA := &entities.Content{
		Content:     []byte{1, 2, 3, 4},
		MetaData:    [][]byte{testByte},
		HashedData:  [][]byte{testByte},
		ContentName: "",
		SplitCount:  0,
		Owner:       "",
		Id:          "",
		UserId:      "",
		ContentId:   "",
	}
	contentB := &entities.Content{
		Content:     []byte{},
		MetaData:    [][]byte{testByte},
		HashedData:  [][]byte{testByte},
		ContentName: "",
		SplitCount:  0,
		Owner:       "",
		Id:          "",
		UserId:      "",
		ContentId:   "",
	}
	chal := &entities.Chal{
		ContentId: "",
		C:         0,
		K1:        []byte{},
		K2:        []byte{},
	}
	_, err := Pg.AuditProofGen(chal, contentA, contentB)
	if err != nil {
		t.Fatal(err)
	}
}
