package entities

type Challenge struct {
	ContentId string `json:"art_id"`
	C         int    `json:"ck"`
	K1        []byte `json:"k1"`
	K2        []byte `json:"k2"`
}

type ChallengeList struct {
	DataList []Challenge `json:"data"`
	Total    int    `json:"total"`
}
