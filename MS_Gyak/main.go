package main

import (
	"net/http"
	//"fmt"
	"github.com/SzMark092/MS_Gyak/internal/HtmlHandlerFuncs"
	"github.com/SzMark092/MS_Gyak/internal/db"
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

	//The http listen code
	http.HandleFunc("/", handler)

	http.ListenAndServe(":8080", nil)
}
