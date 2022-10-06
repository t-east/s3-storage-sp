package entities

type ContentLog struct {
	Owner    string
	Hash     [][]byte
	Provider string
}

type AuditLog struct {
	Chal      *Challenge `json:"chal"`
	Proof     *Proof     `json:"proof"`
	Result    bool       `json:"result"`
	ContentID string     `json:"content_id"`
}
