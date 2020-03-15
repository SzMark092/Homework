package sql_web_handler

import (
	"context"
	"sync"

	"get_table"
	"github.com/SzMark092/MS_Gyak/PSQL_Server/restapi/operations/sql_web_handler"
	"github.com/go-openapi/runtime/middleware"
)

func (p *Pet) create_table(ctx context.Context, params pet.PetCreateParams) middleware.Responder {

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

func (p *Pet) get_table(ctx context.Context, params pet.PetCreateParams) middleware.Responder {

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

func (p *Pet) home_page(ctx context.Context, params pet.PetCreateParams) middleware.Responder {

	wd, _ := os.Getwd()
	http.ServeFile(w, r, wd+"/html/HomePage.html" /*+r.URL.EscapedPath()*/)

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
