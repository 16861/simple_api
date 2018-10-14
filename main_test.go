package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"projects/simple_api/app"
	"projects/simple_api/app/db"
	"testing"
)

func TestRootApp(t *testing.T) {

	response, err := http.Get("http://localhost:8081/")
	if err != nil {
		t.Error("When get root, err: ", err)
	}
	var resp app.RespStatus
	dt, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		t.Error("when read response body, err: ", err)
	}
	err = json.Unmarshal(dt, &resp)
	if err != nil {
		t.Error("when read unmarshal response body, err: ", err)
	}

	if resp.Status != "OK" {
		t.Error("Wrong status code, expect OK actual: ", resp.Status)
	}

}

func TestGetPaintingsApp(t *testing.T) {
	response, err := http.Get("http://localhost:8081/paintings")
	if err != nil {
		t.Error("When get root, err: ", err)
	}
	var resp db.ColectionOfPaintings
	dt, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		t.Error("when read response body, err: ", err)
	}
	err = json.Unmarshal(dt, &resp)
	if err != nil {
		t.Error("when read unmarshal response body, err: ", err)
	}

	if resp.Name != "fist_try" {
		t.Error("Wrong name, actual: ", resp.Name)
	}
}
