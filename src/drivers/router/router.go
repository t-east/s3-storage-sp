package router

import (
	// "database/sql"
	// "fmt"
	"fmt"
	"log"
	"net/http"
	"path"
	"strings"

	// "os"

	// blank import for MySQL driver
	eth "sp/src/drivers/ethereum"
	rdb "sp/src/drivers/rdb"
	auth "sp/src/interfaces/auth"
	"sp/src/interfaces/controllers"

	_ "github.com/go-sql-driver/mysql"
)

func LoadTestDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&entities.User{})
	return db, nil
}

func ServerHandlerPublic(w http.ResponseWriter, r *http.Request) {
	authHandler := auth.NewAuthHandler()
	var head string
	head, r.URL.Path = shiftPath(r.URL.Path)
	switch head {
	case "auth":
		authHandler.Dispatch(w, r)
	default:
		http.Error(w, fmt.Sprintf("method not allowed request. req: %v", r.URL), http.StatusNotFound)
	}
}

// Serve はserverを起動させます．
func Serve() {
	// データベース情報を取得
	db, err := rdb.NewSQLHandler()
	if err != nil {
		log.Fatalf("Can't get DB. %+v", err)
	}

	// パラメータを取得
	param, err := eth.GetParam()
	if err != nil {
		log.Fatalf("Can't get Param from BC. %+v", err)
	}

	private := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var head string
		_, r.URL.Path = shiftPath(r.URL.Path)
		head, r.URL.Path = shiftPath(r.URL.Path)
		switch head {
		case "users":
			uc := controllers.LoadUserController(db)
			uc.Dispatch(w, r)
		case "content":
			cc := controllers.LoadContentController(db, param)
			cc.Dispatch(w, r)
		case "audit":
			ac := controllers.LoadAuditController(db, param)
			ac.Dispatch(w, r)
		default:
			http.Error(w, fmt.Sprintf("method not allowed request. req: %v", r.URL), http.StatusNotFound)
		}
	})

	sm := http.NewServeMux()
	sm.Handle("/api/", auth.JwtMiddleware.Handler(private))
	sm.Handle("/auth/", http.HandlerFunc(ServerHandlerPublic))
	err = http.ListenAndServe(":8080", sm)
	if err != nil {
		log.Fatalf("Listen and serve failed. %+v", err)
	}
}
