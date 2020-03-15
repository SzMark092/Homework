package web

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strconv"

	"github.com/SzMark092/MS_Gyak/internal/db"
)

//html API struct
type RestApi struct {
	PortNum string

	ActSQLhandler *db.Handler
}

//Start listening, get handle functions.
func (h RestApi) StartListen() error {

	http.HandleFunc("/create/", h.createTableResponse)
	http.HandleFunc("/getTable/", h.getTableResponse)
	http.HandleFunc("/", h.getHtmlPage)

	err := http.ListenAndServe(h.PortNum, nil)

	return err
}

//Get a blank data page.
func (h RestApi) getHtmlPage(w http.ResponseWriter, r *http.Request) {
	wd, _ := os.Getwd()
	http.ServeFile(w, r, wd+"/html/HomePage.html" /*+r.URL.EscapedPath()*/)
}

//Get the specified table.
func (h RestApi) getTableResponse(w http.ResponseWriter, r *http.Request) {

	var err error
	var typeCode int
	var dataTable []byte

	typeCode, _ = strconv.Atoi(r.URL.Query().Get("typeCode"))

	dataTable, err = h.makeJson(typeCode)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(dataTable)

}

//Create the specified table.
func (h RestApi) createTableResponse(w http.ResponseWriter, r *http.Request) {

	var err error
	var typeCode int

	typeCode, _ = strconv.Atoi(r.URL.Query().Get("typeCode"))

	switch typeCode {
	case 1:
		h.ActSQLhandler.MakeDataPointDescriptionTable()
	case 2:
		h.ActSQLhandler.MakeDataPointTable()
	case 3:
		h.ActSQLhandler.MakeModuleTable()
	default:
		err = errors.New("Wrong type-code.")
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

//Make a table of data as JSON.
func (h RestApi) makeJson(CodeOfType int) ([]byte, error) {

	var actTypeDataList interface{} = nil

	var err error = nil
	var result []byte

	switch CodeOfType {
	case 1:
		actTypeDataList, err = h.ActSQLhandler.GetDataPointDescriptionTable()
	case 2:
		actTypeDataList, err = h.ActSQLhandler.GetDataPointTable()
	case 3:
		actTypeDataList, err = h.ActSQLhandler.GetModuleTable()
	default:
		err = errors.New("Wrong type-code.")
	}
	if err != nil {
		return result, err
	}

	result, err = json.Marshal(actTypeDataList)

	return result, err

}
