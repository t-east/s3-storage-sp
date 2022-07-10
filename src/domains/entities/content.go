package entities

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type ContentIn struct {
	Address  string   `json:"address"`
	Content  Point    `json:"content"`
	MetaData [][]byte `json:"meta_data"`
}

type ContentInForUser struct {
	Content Point  `json:"content"`
	PrivKey string `json:"priv_key"`
	Address string `json:"address"`
}

type ContentInDB struct {
	ID       string   `json:"id"`
	Content  Point    `json:"content"`
	MetaData [][]byte `json:"metadata"`
}

type MetaDataInDB struct {
	ID        string `json:"id"`
	ContentID uint   `json:"content_id"`
	MetaData  []byte `json:"metadata"`
}

type Content struct {
	ID       string   `json:"id"`
	Address  string   `json:"address"`
	Content  Point    `json:"content"`
	MetaData [][]byte `json:"metadata"`
	HashData [][]byte `json:"hashdata"`
}

type ContentLog struct {
	Owner    string
	Hash     [][]byte
	Provider string
}

type Receipt struct {
	ID       string   `json:"id"`
	Content  Point    `json:"content"`
	MetaData [][]byte `json:"metadata"`
	HashData [][]byte `json:"hashdata"`
}

func NewContent() *ContentIn {
	return &ContentIn{}
}

type ContentInStorage struct {
	Id       string   `json:"id"`
	File     []byte   `json:"file"`
	MetaData []string `json:"meta_data"`
	FileName string   `json:"name"`
}

type ContentInBlockChain struct {
	HashedData [][]byte `json:"hashed_data"`
	ContentId  string   `json:"content_id"`
	Owner      string   `json:"owner"`
}

type Key struct {
	PubKey  string `json:"pubkey"`
	PrivKey string `json:"privkey"`
}

type Log struct {
	AuditLog   []*AuditLog            `json:"audit_log"`
	ContentLog []*ContentInBlockChain `json:"content_log"`
}
