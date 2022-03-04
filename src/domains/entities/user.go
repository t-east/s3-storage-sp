package entities

type User struct {
	ID      string `json:"id"`
	Address string `json:"address"`
	PubKey  string `json:"pub_key"`
	PrivKey string `json:"priv_key"`
}

type UserRequest struct {
	Address  string `json:"address"`
	Password string `json:"password"`
}
