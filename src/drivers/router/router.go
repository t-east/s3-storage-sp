package router

import (
	// "database/sql"
	// "fmt"
	"log"
	"net/http"

	// "os"

	// blank import for MySQL driver
	eth "sp/src/drivers/ethereum"
	rdb "sp/src/drivers/rdb"
	"sp/src/interfaces/controllers"

	_ "github.com/go-sql-driver/mysql"
)

// Serve はserverを起動させます．
func Serve(addr string) {
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

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Listen and serve failed. %+v", err)
	}
}
