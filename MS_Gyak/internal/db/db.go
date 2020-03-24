package db

import (
	"errors"
	"fmt"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"

	models "github.com/SzMark092/MS_Gyak/PSQL_Server/models"
)

/*
https://godoc.org/github.com/go-pg/pg
https://github.com/go-pg/pg/wiki/Writing-Queries
https://github.com/go-pg/pg/wiki/FAQ#how-to-view-queries-this-library-generates
*/

//Make the connection, and give itt back to the caller.
func MakeConnection(name, password, serverName, database string) *pg.DB {
	conn := pg.Connect(&pg.Options{
		User:     name,
		Password: password,
		Addr:     serverName,
		Database: database,
	})

	if _, err := conn.Exec("SELECT 1"); err != nil {
		panic(err)
	}

	return conn
}

type PSQLServerConfig struct {
	Name       string
	Password   string
	ServerName string
	Database   string
}

type Handler struct {
	conn      *pg.DB
	tempTable bool
}

func NewHandler(actConfig PSQLServerConfig, tempTable bool, createAllTable bool) (*Handler, error) {
	handler := &Handler{
		tempTable: tempTable,
	}

	handler.conn = MakeConnection(actConfig.Name, actConfig.Password, actConfig.ServerName, actConfig.Database)

	if createAllTable {

		if err := handler.makeTable(&models.DataPointDescription{}); err != nil {
			return nil, err
		}

		if err := handler.makeTable(&models.Module{}); err != nil {
			return nil, err
		}

		if err := handler.makeTable(&models.DataPoint{}); err != nil {
			return nil, err
		}

	}
	return handler, nil
}

//Get the model type connected to the given code.
func getModelType(CodeOfType int) (interface{}, interface{}, error) {

	var modelType interface{}
	var modelTypeList interface{}
	var err error

	switch CodeOfType {

	case 1:
		modelType = &models.DataPointDescription{}
		modelTypeList = &[]models.DataPointDescription{}
	case 2:
		modelType = &models.DataPoint{}
		modelTypeList = &[]models.DataPoint{}
	case 3:
		modelType = &models.Module{}
	default:
		err = errors.New("Wrong type-code.")
	}

	return modelType, modelTypeList, err
}

//Make table from the given struct.
func (h Handler) makeTable(actModelStruct interface{}) error {
	return h.conn.CreateTable(actModelStruct, &orm.CreateTableOptions{
		Temp:          h.tempTable,
		FKConstraints: true,
		IfNotExists:   true,
	})
}

//Close the connection.
func (h Handler) Close() error {
	return h.conn.Close()
}

//Place in data with the given type.
func (h Handler) ModelInsert(actModelStruct interface{}) (interface{}, error) {
	result, err := h.conn.Model(actModelStruct).Returning("*").Insert()
	if err == nil && result.RowsReturned() != 1 {
		err = fmt.Errorf("invalid result: %d", result.RowsReturned())
	}
	return actModelStruct, err
}

//Get the DataPointDescription table.
func (h Handler) GetDataPointDescriptionTable() ([]*models.DataPointDescription, error) {

	var rows []*models.DataPointDescription

	count, err := h.conn.Model(&models.DataPointDescription{}).Count()
	if err != nil {
		return nil, err
	}

	err = h.conn.Model(&rows).Limit(count).Select()
	if err != nil {
		return nil, err
	}

	return rows, err

}

//Get the DataPoint table
func (h Handler) GetDataPointTable() ([]*models.DataPoint, error) {

	var rows []*models.DataPoint

	count, err := h.conn.Model(&models.DataPoint{}).Count()
	if err != nil {
		return nil, err
	}

	err = h.conn.Model(&rows).Limit(count).Select()
	if err != nil {
		return nil, err
	}

	return rows, err

}

//Get the DataPoint table
func (h Handler) GetModuleTable() ([]*models.Module, error) {

	var rows []*models.Module

	count, err := h.conn.Model(&models.Module{}).Count()
	if err != nil {
		return nil, err
	}

	err = h.conn.Model(&rows).Limit(count).Select()
	if err != nil {
		return nil, err
	}

	return rows, err

}

//Make DataPointDescription table.
func (h Handler) MakeDataPointDescriptionTable() error {
	err := h.makeTable(&models.DataPointDescription{})
	return err
}

//Make DataPoint table.
func (h Handler) MakeDataPointTable() error {
	err := h.makeTable(&models.DataPoint{})
	return err
}

//Make Module table.
func (h Handler) MakeModuleTable() error {
	err := h.makeTable(&models.Module{})
	return err
}
