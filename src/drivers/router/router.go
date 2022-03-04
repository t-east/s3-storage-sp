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

	// コントローラの準備
	_ = controllers.LoadUserController(db, param)
	_ = controllers.LoadContentController(db, param)

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Listen and serve failed. %+v", err)
	}
}
