package router

import (
	// "database/sql"
	// "fmt"
	"fmt"
	"log"
	"net/http"

	// "os"

	// blank import for MySQL driver
	"sp/src/core"
	"sp/src/domains/entities"

	// rdb "sp/src/drivers/rdb"
	"sp/src/interfaces/controllers"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func LoadTestDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&entities.User{})
	db.AutoMigrate(&entities.Receipt{})
	return db, nil
}

// func ServerHandlerPublic(w http.ResponseWriter, r *http.Request) {
// 	authHandler := auth.NewAuthHandler()
// 	var head string
// 	head, r.URL.Path = core.ShiftPath(r.URL.Path)
// 	switch head {
// 	case "auth":
// 		authHandler.Dispatch(w, r)
// 	default:
// 		http.Error(w, fmt.Sprintf("method not allowed request. req: %v", r.URL), http.StatusNotFound)
// 	}
// }

func allowOptionsMiddleware(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return nil
	}
	return nil
}

func ServerHandlerPublic(w http.ResponseWriter, r *http.Request) {
	db, err := LoadTestDB()
	if err != nil {
		log.Fatalf("Can't get DB. %+v", err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	allowOptionsMiddleware(w, r)
	var head string
	_, r.URL.Path = core.ShiftPath(r.URL.Path)
	head, r.URL.Path = core.ShiftPath(r.URL.Path)
	switch head {
	case "content":
		uc := controllers.LoadContentController(db)
		uc.Dispatch(w, r)
	default:
		http.Error(w, fmt.Sprintf("method not allowed request. req: %v", r.URL), http.StatusNotFound)
	}
}

// Serve はserverを起動させます．
func Serve() {
	sm := http.NewServeMux()
	sm.Handle("/api/", http.HandlerFunc(ServerHandlerPublic))
	err := http.ListenAndServe(":4001", sm)
	if err != nil {
		log.Fatalf("Listen and serve failed. %+v", err)
	}
}

// // Serve はserverを起動させます．
// func Serve() {
// 	// データベース情報を取得
// 	db, err := LoadTestDB()
// 	if err != nil {
// 		log.Fatalf("Can't get DB. %+v", err)
// 	}

// 	log.Print(db)
// 	// パラメータを取得
// 	param := &contracts.Param{
// 		Paring: "a",
// 		G:      []byte{1},
// 		U:      []byte{1},
// 	}

// 	private := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		var head string
// 		_, r.URL.Path = core.ShiftPath(r.URL.Path)
// 		head, r.URL.Path = core.ShiftPath(r.URL.Path)
// 		switch head {
// 		case "users":
// 			uc := controllers.LoadUserController(db)
// 			uc.Dispatch(w, r)
// 		case "content":
// 			cc := controllers.LoadContentController(db)
// 			cc.Dispatch(w, r)
// 		case "audit":
// 			ac := controllers.LoadAuditController(db, param)
// 			ac.Dispatch(w, r)
// 		default:
// 			http.Error(w, fmt.Sprintf("method not allowed request. req: %v", r.URL), http.StatusNotFound)
// 		}
// 	})

// 	sm := http.NewServeMux()
// 	sm.Handle("/api/", private)
// 	sm.Handle("/auth/", http.HandlerFunc(ServerHandlerPublic))
// 	err = http.ListenAndServe(":4000", sm)
// 	if err != nil {
// 		log.Fatalf("Listen and serve failed. %+v", err)
// 	}
// }
