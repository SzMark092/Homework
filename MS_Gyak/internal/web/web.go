package web

import (
	"context"
	"errors"

	/*"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"*/

	"github.com/SzMark092/MS_Gyak/PSQL_Server/models"
	"github.com/SzMark092/MS_Gyak/PSQL_Server/restapi/operations/sql_web_handler"
	"github.com/SzMark092/MS_Gyak/internal/db"
	"github.com/go-openapi/runtime/middleware"
)

type SqlWebHandler struct {
	ActSQLhandler *db.Handler
}

func New() (SqlWebHandler, error) {

	var actHandler SqlWebHandler
	var err error

	actConfig := &db.PSQLServerConfig{
		Name:       "postgres",
		Password:   "1234",
		ServerName: "localhost:5432",
		Database:   "prod",
	}

	actHandler.ActSQLhandler, err = db.NewHandler(*actConfig, false, false)

	return actHandler, err

}

func (h *SqlWebHandler) CreateTable(ctx context.Context, params sql_web_handler.CreateTableParams) middleware.Responder {

	var err error
	var typeCode int64

	typeCode = params.TableType

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

		return sql_web_handler.NewCreateTableMethodNotAllowed()

	}

	return sql_web_handler.NewCreateTableOK()

}

func (h *SqlWebHandler) GetDataPointDescriptionTable(ctx context.Context, params sql_web_handler.GetDataPointDescriptionTableParams) middleware.Responder {

	actTypeDataList, err := h.ActSQLhandler.GetDataPointDescriptionTable()
	var result = make([]*models.DataPointDescription, len(actTypeDataList))

	result = actTypeDataList

	if err != nil || len(actTypeDataList) == 0 {
		return sql_web_handler.NewGetDataPointDescriptionTableNotFound()
	}

	return sql_web_handler.NewGetDataPointDescriptionTableOK().WithPayload(result)

}

func (h *SqlWebHandler) GetDataPointTable(ctx context.Context, params sql_web_handler.GetDataPointTableParams) middleware.Responder {

	actTypeDataList, err := h.ActSQLhandler.GetDataPointTable()
	var result = make([]*models.DataPoint, len(actTypeDataList))

	result = actTypeDataList

	if err != nil || len(actTypeDataList) == 0 {
		return sql_web_handler.NewGetDataPointTableNotFound()
	}

	return sql_web_handler.NewGetDataPointTableOK().WithPayload(result)

}

func (h *SqlWebHandler) GetModule(ctx context.Context, params sql_web_handler.GetModuleParams) middleware.Responder {

	actTypeDataList, err := h.ActSQLhandler.GetModuleTable()
	var result = make([]*models.Module, len(actTypeDataList))

	result = actTypeDataList

	if err != nil || len(actTypeDataList) == 0 {

		return sql_web_handler.NewGetModuleNotFound()

	}

	return sql_web_handler.NewGetModuleOK().WithPayload(result)

}

//Make a table of data as JSON.
/*func (h *SqlWebHandler) choose_table(CodeOfType int64) ([]interface{}, error) {

	var actTypeDataList []interface{} = nil

	var err error = nil

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
		return actTypeDataList, err
	}

	return actTypeDataList, err

}*/
