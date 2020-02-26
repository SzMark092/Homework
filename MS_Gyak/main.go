package main

import (

	//"fmt"
	"github.com/SzMark092/MS_Gyak/internal/db"
	"github.com/SzMark092/MS_Gyak/internal/web"
)

func main() {
	conn := db.MakeConnection("postgres", "1234", "localhost:5432", "prod")

	var err error
	var GlobalHandler *db.Handler

	GlobalHandler, err = db.NewHandler(conn, false)

	if err != nil {
		panic(err)
	}

	defer GlobalHandler.Close()

	restApi := web.RestApi{
		PortNum:       ":8080",
		ActSQLhandler: GlobalHandler,
	}
	restApi.StartListen()
}
