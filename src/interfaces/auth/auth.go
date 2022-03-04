package handler

import (
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/form3tech-oss/jwt-go"
)

type RequestAuth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthHandler struct {
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

var secretKey = "SECRET_KEY" // TODO: SECRET_KEYを別の値に

// AccessTokenの取得を行うHandler TODO: refreshトークンも？
func (ah *AuthHandler) Post(w http.ResponseWriter, r *http.Request) (string, error) {
	// 推奨のjwt.StandardClaimsを利用する方法
	claims := &jwt.StandardClaims{
		Subject: "123",
		Issuer:  "123",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 電子署名 TODO: 適当なKeyに置き換える
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (ah *AuthHandler) Dispatch(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		ah.Post(w, r)
	default:
		http.NotFound(w, r)
	}
}

var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		// TODO: 適当なKeyに置き換える
		return []byte(secretKey), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})
