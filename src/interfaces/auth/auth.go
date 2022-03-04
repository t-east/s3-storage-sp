package auth

import (
	"encoding/json"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/form3tech-oss/jwt-go"
)

type RequestAuth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type responseToken struct {
	AccessToken string `json:"access_token"`
}

type AuthHandler struct {
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

var secretKey = "SECRET_KEY" // TODO: SECRET_KEYを別の値に

// AccessTokenの取得を行うHandler TODO: refreshトークンも？
func (ah *AuthHandler) Post(w http.ResponseWriter, r *http.Request) {
	// TODO: パスワード間違ってても処理が通る。。。
	// $ curl -X POST -H "Content-Type: application/json" -d '{"email":"a@example.com", "password":"test"}' localhost:8080/auth
	// {"access_token":"eyJhbG..."}
	// bodyBytes, err := ioutil.ReadAll(r.Body)
	// defer r.Body.Close()
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	// 	return
	// }
	// body := RequestAuth{}
	// if err = json.Unmarshal(bodyBytes, &body); err != nil {
	// 	http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	// 	return
	// }
	// TODO ユーザ呼び出し
	// foundUser, err2 := ah.userUsecase.FindByEmail(body.Email)
	// if err2 != nil {
	// 	http.Error(w, err2.Error(), http.StatusNotFound)
	// 	return
	// }

	// 推奨のjwt.StandardClaimsを利用する方法
	claims := &jwt.StandardClaims{
		Subject: "sdf",
		Issuer:  "sdf",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 電子署名 TODO: 適当なKeyに置き換える
	tokenString, err3 := token.SignedString([]byte(secretKey))
	if err3 != nil {
		http.Error(w, err3.Error(), http.StatusInternalServerError)
		return
	}

	tokenResp := responseToken{AccessToken: tokenString}
	res, err4 := json.Marshal(tokenResp)
	if err4 != nil {
		http.Error(w, err4.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func (ah *AuthHandler) Dispatch(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		ah.Post(w, r)
	default:
		http.NotFound(w, r)
	}
}

var JwtMiddleware = jwtmiddleware.New(
	jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			// TODO: 適当なKeyに置き換える
			return []byte(secretKey), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	},
)
