package db

import (
	"fmt"
	"math/rand"

	"testing"
	//"github.com/stretchr/testify/assert"
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

func makeTestDb() *Handler {

	conn := MakeConnection("postgres", "1234", "localhost:5432", "TestDb")

	handler, err := NewHandler(conn, true, false)
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

	fmt.Println("Make DataPointDescription table.")

	var actData = DataPointDescription{1, "fgfd", "sfsdf", 5, 10}

	if err := handler.MakeTable(&actData); err != nil {
		fmt.Println("Making of DataPointDescription table failed.")
		return
	}

	if err = handler.ModelInsert(actData); err != nil {
		fmt.Println("Making of DataPointDescription table failed.")
		return
	}

}

func TestGetAllDataOfType(*testing.T) {

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
