package app

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"projects/simple_api/app/db"
)

//Route contains path method and func
type Route struct {
	Path   string
	Method string
	Fn     HTTPFunc
}

//RespStatus contains return status, default is OK
type RespStatus struct {
	Status string `json:"status"`
}

//Routes isarray for routers
type Routes []Route

func getDB() *db.DB {
	return &db.DB{
		User:           "api_test",
		Password:       "api_test1",
		Path:           "ds017185.mlab.com:17185",
		DBName:         "mdb_test",
		CollectionName: "paintings",
	}
}

//Root func
func Root(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "Application/json")
	data, err := json.Marshal(RespStatus{"OK"})
	if err != nil {
		log.Println("error when unmarshal status struct, err: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

//GetPictures returns pictures from mongo db
func GetPictures(w http.ResponseWriter, r *http.Request) {
	db := getDB()
	paintings := db.GetPantings()
	w.Header().Add("Content-type", "Application/json")
	payload, err := json.Marshal(paintings)
	if err != nil {
		log.Print("Can't unmarshal paintings, err: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(payload)
}

//AddPainting add new painting to db
func AddPainting(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		log.Println("Error when read body, err: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var painting db.Painting
	err = json.Unmarshal(data, &painting)
	if err != nil {
		log.Println("Error when unmarshal body, err: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = getDB().AddPainting(painting)
	if err != nil {
		log.Println("Error when updating db, err: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

//GetRoutes returns all possible routes
func GetRoutes() Routes {
	routers := Routes{}

	routers = append(routers, Route{"/", "GET", Root})
	routers = append(routers, Route{"/paintings", "GET", GetPictures})
	routers = append(routers, Route{"/paintings", "POST", AddPainting})

	return routers
}
