package config

var (
	// "development", "test", "prod"
	env string
	// JWTトークンの署名
	sigKey string
)

func init() {
	// TODO 将来的には.envなどから読み込むが一旦これで固定
	env = "development"
	sigKey = "XXX"
}

func IsDevelopment() bool {
	return env == "development"
}

func IsTest() bool {
	return env == "test"
}

func SigKey() string {
	return sigKey
}
