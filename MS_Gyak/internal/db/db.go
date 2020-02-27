package db

import (
	"errors"
	"fmt"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
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

type Handler struct {
	conn      *pg.DB
	tempTable bool
}

func NewHandler(conn *pg.DB, tempTable bool, createAllTable bool) (*Handler, error) {
	handler := &Handler{
		conn:      conn,
		tempTable: tempTable,
	}

	if createAllTable {

		if err := handler.MakeTable(&DataPointDescription{}); err != nil {
			return nil, err
		}

		if err := handler.MakeTable(&Module{}); err != nil {
			return nil, err
		}

		if err := handler.MakeTable(&DataPoint{}); err != nil {
			return nil, err
		}

	}
	return handler, nil
}

//Get the model type connected to the given code.
func GetModelType(CodeOfType int) (interface{}, interface{}, error) {

	var modelType interface{}
	var modelTypeList interface{}
	var err error

	switch CodeOfType {

	case 1:
		modelType = &DataPointDescription{}
		modelTypeList = &[]DataPointDescription{}
	case 2:
		modelType = &DataPoint{}
		modelTypeList = &[]DataPoint{}
	case 3:
		modelType = &Module{}
	default:
		err = errors.New("Wrong type-code.")
	}

	return modelType, modelTypeList, err
}

//Make table from the given struct.
func (h Handler) MakeTable(actModelStruct interface{}) error {
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

//Select all data of a specific type.
func (h Handler) SelectAllData(actModelStructToSelect interface{}) ([]interface{}, error) {
	var rows []interface{}

	count, err := h.conn.Model(actModelStructToSelect).Count()
	if err != nil {
		return nil, err
	}

	err = h.conn.Model(&rows).Limit(count).Select()
	if err != nil {
		return nil, err
	}

	return rows, nil

}
