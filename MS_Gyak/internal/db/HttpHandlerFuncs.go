package db

import (
	"encoding/json"
	"net/http"
	"strconv"
)

//html API struct
type html_API struct {
	PortNum string

	ActSQLhandler Handler
}

//Start listening, get handle functions.
func (h html_API) startListen() error {

	http.HandleFunc("/create/", h.createTableResponse)
	http.HandleFunc("/getTable/", h.getTableResponse)
	http.HandleFunc("/getDataPage/", h.getDataPage)

	err := http.ListenAndServe(h.PortNum, nil)

	return err
}

//Get a blank data page.
func (h html_API) getDataPage(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "DataPage.html")

}

//Get the specified table.
func (h html_API) getTableResponse(w http.ResponseWriter, r *http.Request) {

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
func (h html_API) createTableResponse(w http.ResponseWriter, r *http.Request) {

	var err error
	var typeCode int
	var actTypeData interface{}

	r.ParseForm()

	typeCode, _ = strconv.Atoi(r.Form.Get("typeCode"))

	actTypeData, err = getModelType(typeCode)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.ActSQLhandler.makeTable(actTypeData)

}

//Make a table of data as JSON.
func (h html_API) makeJson(CodeOfType int) ([]byte, error) {

	var actTypeData interface{} = nil
	var actTypeDataList []interface{} = nil

	var err []error = nil //Store errors.
	var result []byte

	actTypeData, err[0] = getModelType(CodeOfType)

	actTypeDataList, err[0] = h.ActSQLhandler.SelectAllData(actTypeData)

	result, err[1] = json.Marshal(actTypeDataList)

	if err[0] != nil {
		return result, err[0]
	}

	return result, err[1]

}
