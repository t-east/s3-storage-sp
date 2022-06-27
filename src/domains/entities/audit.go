package entities

type ArtLog struct {
	HashedData []string `json:"hashed_data"`
	Owner      string   `json:"owner"`
}

type Chal struct {
	ContentId string `json:"art_id"`
	C         int    `json:"ck"`
	K1        string `json:"k1"`
	K2        string `json:"k2"`
}

type Proof struct {
	Myu       string `json:"myu"`
	Gamma     string `json:"gamma"`
	ContentId string `json:"content_id"`
}

type Proofs struct {
	DataList []Proof `json:"proofs"`
	Total    int     `json:"total"`
}

type Chals struct {
	DataList []Chal `json:"data"`
	Total    int    `json:"total"`
}
