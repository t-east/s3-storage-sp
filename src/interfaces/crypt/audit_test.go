package crypt

import (
	"encoding/json"
	"io/ioutil"
	"sp/src/core"
	"sp/src/domains/entities"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAudit(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		orgID    string
		wantCode int
		wantErr  bool
	}{
		{
			name:    "Success",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			param, key, err := core.CreateParamMock()
			if err != nil {
				t.Fatal(err)
			}

			contentIn := &entities.ContentInForUser{
				Content: entities.Point{
					X: 1,
					Y: 2,
				},
				PrivKey: string(key.PrivKey),
				Address: "",
			}

			ac := &AuditCrypt{
				Param: param,
			}

			content, err := MakeMetaData(contentIn, param)
			if err != nil {
				t.Fatal(err)
			}

			content.ID = core.MakeULID()

			file, _ := json.MarshalIndent(content, "", " ")
			_ = ioutil.WriteFile("content.json", file, 0644)

			receipt := &entities.Receipt{
				ID:       content.ID,
				Content:  content.Content,
				MetaData: content.MetaData,
				HashData: content.HashData,
			}

			cInB := &entities.ContentInBlockChain{
				HashedData: content.HashData,
				ContentId:  content.ID,
				Owner:      content.Address,
			}
			chal := AuditChallen(param)
			chal.ContentId = content.ID
			file, _ = json.MarshalIndent(chal, "", " ")
			_ = ioutil.WriteFile("chal.json", file, 0644)
			proof, err := ac.AuditProofGen(chal, receipt, cInB)
			if err != nil {
				t.Fatal(err)
			}
			proof.ContentId = content.ID
			file, _ = json.MarshalIndent(proof, "", " ")
			_ = ioutil.WriteFile("proof.json", file, 0644)

			log := &entities.Log{
				AuditLog: []*entities.AuditLog{
					{
						Chal:      chal,
						Proof:     proof,
						Result:    false,
						ContentID: cInB.ContentId,
					},
				},
				ContentLog: []*entities.ContentInBlockChain{
					{
						HashedData: content.HashData,
						ContentId:  cInB.ContentId,
						Owner:      content.Address,
					},
				},
			}
			file, _ = json.MarshalIndent(log, "", " ")
			_ = ioutil.WriteFile("log.json", file, 0644)

			left, right, err := AuditVerify(param, key.PubKey, content, proof, chal)
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(left, right) {
				t.Errorf("Verify diff = %v, %v, %v", cmp.Diff(left, right), left, right)
			}
		})
	}
}
