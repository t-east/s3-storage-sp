package crypt

import (
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

			chal := AuditChallen(param)

			proof, err := ac.AuditProofGen(chal, content, content)
			if err != nil {
				t.Fatal(err)
			}

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
