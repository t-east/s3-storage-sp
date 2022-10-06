package entities

type Proof struct {
	Myu       []byte `json:"myu"`
	Gamma     []byte `json:"gamma"`
	ContentId string `json:"content_id"`
}

type ProofList struct {
	DataList []Proof `json:"proofs"`
	Total    int     `json:"total"`
}
