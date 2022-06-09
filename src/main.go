package main

import (
	"sp/src/domains/entities"
	"sp/src/drivers/router"
	"sp/src/interfaces/controllers"
)

func realMain() {
	cc := *controllers.LoadContentController(&entities.Param{})
	ac := *controllers.LoadAuditController(&entities.Param{})

	e := router.NewServer(cc, ac)
	e.Logger.Fatal(e.Start(":4001"))
}

func main() {
	realMain()
}
