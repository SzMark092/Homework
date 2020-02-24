package db

import (
	"fmt"

	"testing"

	"github.com/stretchr/testify/assert"
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

func makeTestDb() *Handler {
	/*
		conn := MakeConnection("postgres", "1234", "localhost:5432", "")
		if _, err := conn.Exec("CREATE DATABASE TestDb;"); err != nil {
			panic(err)
		}
	*/
	conn := MakeConnection("postgres", "1234", "localhost:5432", "TestDb")

	handler, err := NewHandler(conn, true)
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
