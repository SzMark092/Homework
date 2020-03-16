package sql_web_handler

import (
	"encoding/json"
	"io/ioutil"
	//"fmt"
	"github.com/SzMark092/MS_Gyak/internal/db"
	//"math/rand"
	"net/http"
	//"strconv"
	//"github.com/stretchr/testify/assert"
	"testing"
)

var restApi RestApi

func buildTestAPI() {
	conn := db.MakeConnection("postgres", "1234", "localhost:5432", "prod")

	var err error
	var GlobalHandler *db.Handler

	GlobalHandler, err = db.NewHandler(conn, false, false)

	if err != nil {
		panic(err)
	}

	defer GlobalHandler.Close()

	restApi = RestApi{
		PortNum:       ":8080",
		ActSQLhandler: GlobalHandler,
	}
	restApi.StartListen()
}

func TestMakeTableReq(t *testing.T) {

	go buildTestAPI()

	t.Log("Test API is ready.")

	client := &http.Client{}

	resp, err := client.Get("http://localhost:8080/create?typeCode=1")
	resp, err = client.Get("http://localhost:8080/create?typeCode=2")
	resp, err = client.Get("http://localhost:8080/create?typeCode=3")

	if err != nil {
		t.Log(err.Error())
	}

	t.Log(resp)

}

func TestGetTableReq(t *testing.T) {

	go buildTestAPI()

	t.Log("Test API is ready.")

	client := &http.Client{}

	resp, err := client.Get("http://localhost:8080/getTable?typeCode=1")

	actBody, err := ioutil.ReadAll(resp.Body)

	var result []db.DataPointDescription

	json.Unmarshal(actBody, &result)

	if err != nil {
		t.Log(err.Error())
	}

	t.Log(result)

}
