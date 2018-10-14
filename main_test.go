package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"projects/simple_api/app"
	"projects/simple_api/app/db"
	"testing"
)

func TestRootApp(t *testing.T) {
	addr := os.Getenv("API_ADDR")
	port := os.Getenv("API_PORT")
	if addr == "" || port == "" {
		t.Fatal("API_PORT or API_ADDR not set!")
	}

	response, err := http.Get(addr + ":" + port)
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
		t.Log("Struct: ", string(dt))
		t.Error("when read unmarshal response body, err: ", err)
	}

	if resp.Status != "OK" {
		t.Error("Wrong status code, expect OK actual: ", resp.Status)
	}

}

func TestGetPaintingsApp(t *testing.T) {
	addr := os.Getenv("API_ADDR")
	port := os.Getenv("API_PORT")
	if addr == "" || port == "" {
		t.Fatal("API_PORT or API_ADDR not set!")
	}

	response, err := http.Get(addr + ":" + port + "/paintings")
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
