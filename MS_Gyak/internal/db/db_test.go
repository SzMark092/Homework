package db

import (
	"fmt"
	"math/rand"
	"reflect"
	"strconv"

	//"github.com/stretchr/testify/assert"
	"testing"
)

/*
https://docs.docker.com/docker-for-windows/install/
https://hub.docker.com/_/postgres
Running Postgres server in Docker container:
	docker run -it --rm -p 5432:5432 --name pgserver --hostname=pgserver -e POSTGRES_PASSWORD=1234 postgres
PG GUI on localhost:8888
	docker run --rm -p 8888:8080 --link pgserver -e ADMINER_DEFAULT_SERVER=PostgreSQL adminer
		System: PostgreSQL
		Server: pgserver
		Username: postgres
		Password: 1234
		Database: postgres
*/

func randomString(length int) string {

	var result string

	for i := 0; i < length; i++ {

		result += string(rune(32 + int32(32*rand.Float32())))

	}

	return result
}

func randomDataPDesc(numberOfData int) []DataPointDescription {

	var listOfStructs = make([]DataPointDescription, numberOfData)

	for i := 0; i < numberOfData; i++ {

		var actStruct DataPointDescription

		actStruct.ID = int64(rand.Int())
		actStruct.Name = randomString(10)
		actStruct.Title = randomString(10)
		actStruct.Min = rand.Int()
		actStruct.Max = rand.Int()

		listOfStructs[i] = actStruct

	}

	return listOfStructs
}

func makeTestDb() *Handler {

	conn := MakeConnection("postgres", "1234", "localhost:5432", "prod")

	handler, err := NewHandler(conn, false, false)
	if err != nil {
		panic(err)
	}

	fmt.Println("Database connection is ready.")

	return handler
}

func TestMakeConn(*testing.T) {
	handler := makeTestDb()
	defer handler.Close() //nolint:errcheck,gocritic

	println("Try to close connection.")
}

func TestInsertSomeData(*testing.T) {

	handler := makeTestDb()
	var err error

	fmt.Println("Make DataPointDescription table.")

	var actData = DataPointDescription{1, "fgfd", "sfsdf", 5, 10}

	if err = handler.MakeTable(&actData); err != nil {
		fmt.Println("Making of DataPointDescription table failed.")
		return
	}

	if _, err = handler.ModelInsert(&actData); err != nil {
		fmt.Println("Making of DataPointDescription table failed. Error:" + err.Error())
		return
	}

}

func TestGetAllDataOfType(t *testing.T) {

	handler := makeTestDb()
	var result []interface{}
	var err error

	if result, err = handler.SelectAllData(DataPointDescription{}); err != nil {
		t.Logf("Failed to get data from table.")
		return
	}

	for i := 0; i < len(result); i++ {
		t.Logf(strconv.Itoa(i)+". row: "+"%+v", result[i])
	}

}

func TestGetModeType(*testing.T) {
	var results [3]interface{}
	var err error

	for i := 0; i < 3; i++ {

		results[i], _, err = GetModelType(1)
		if err != nil {

			println(err.Error())

		} else {

			println("The type of the actual model:" + reflect.TypeOf(results[i]))

		}

		err = nil
	}

}

/*
func TestDataPointDescriptionInsert(t *testing.T) {
	handler := makeTestDb()
	defer handler.Close()

	dpd, err := handler.DataPointDescriptionInsert(
		"NAME", "TITLE",
		100, 1,
	)

	assert.Nil(t, err, "INSERT")
	assert.NotZero(t, dpd.ID, "INSERT")
}
*/
