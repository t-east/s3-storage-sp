package entities

type SampleData struct {
	Name     string `json:"name"`
	Length   int    `json:"length"`
	Location int    `json:"location"`
}

type ContentIn struct {
	Address  string     `json:"address"`
	Content  SampleData `json:"content"`
	MetaData [][]byte   `json:"meta_data"`
}

type ContentInDB struct {
	ID       string     `json:"id"`
	Content  SampleData `json:"content"`
	MetaData [][]byte   `json:"metadata"`
}

type MetaDataInDB struct {
	ID        string `json:"id"`
	ContentID uint   `json:"content_id"`
	MetaData  []byte `json:"metadata"`
}

type Content struct {
	ID       string     `json:"id"`
	Address  string     `json:"address"`
	Content  SampleData `json:"content"`
	MetaData [][]byte   `json:"metadata"`
	HashData []string   `json:"hashdata"`
}

type ContentLog struct {
	Owner    string
	Hash     []string
	Provider string
}

type Receipt struct {
	ID       string     `json:"id"`
	Content  SampleData `json:"content"`
	MetaData [][]byte   `json:"metadata"`
	HashData []string   `json:"hashdata"`
}

func NewContent() *ContentIn {
	return &ContentIn{}
}

type ContentInStorage struct {
	Id       string   `json:"id"`
	File     []byte   `json:"file"`
	MetaData [][]byte `json:"meta_data"`
	FileName string   `json:"name"`
}

type ContentInBlockChain struct {
	MetaData   [][]byte `json:"meta_data"`
	HashedData [][]byte `json:"hashed_data"`
	FileName   string   `json:"name"`
	SplitCount int      `json:"split_count"`
	Owner      string   `json:"owner"`
	ContentId  string   `json:"content_id"`
}

type Key struct {
	PubKey  []byte `json:"pubkey"`
	PrivKey []byte `json:"privkey"`
}
