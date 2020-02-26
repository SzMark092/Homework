package web

import (
	"encoding/json"
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
	http.ServeFile(w, r, wd+"/"+r.URL.EscapedPath())
}

//Get the specified table.
func (h RestApi) getTableResponse(w http.ResponseWriter, r *http.Request) {

	var err error
	var typeCode int
	var dataTable []byte

	r.ParseForm()
	typeCode, _ = strconv.Atoi(r.Form.Get("typeCode"))
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
	var actTypeData interface{}

	r.ParseForm()

	typeCode, _ = strconv.Atoi(r.Form.Get("typeCode"))

	actTypeData, err = db.GetModelType(typeCode)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.ActSQLhandler.MakeTable(actTypeData)

}

//Make a table of data as JSON.
func (h RestApi) makeJson(CodeOfType int) ([]byte, error) {

	var actTypeData interface{} = nil
	var actTypeDataList []interface{} = nil

	var err []error = nil //Store errors.
	var result []byte

	actTypeData, err[0] = db.GetModelType(CodeOfType)

	actTypeDataList, err[0] = h.ActSQLhandler.SelectAllData(actTypeData)

	result, err[1] = json.Marshal(actTypeDataList)

	if err[0] != nil {
		return result, err[0]
	}

	return result, err[1]

}
