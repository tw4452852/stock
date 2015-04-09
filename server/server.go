package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var r = mux.NewRouter()

func init() {
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/query", queryHandler).
		Queries("kind", "{kind:s[h|z]}").
		Queries("id", "{id:[0-9]+}")
}

func homeHandler(res http.ResponseWriter, req *http.Request) {
	log.Println("homeHandler")
	queryAll()
	if e := homeTem.Execute(res, nil); e != nil {
		log.Printf("homeHandler: %s\n", e)
	}
}

func queryHandler(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, exist := vars["id"]
	if !exist {
		log.Println("queryHandler: id is null")
		return
	}
	key := vars["kind"] + id
	log.Printf("queryHandler: key is %s\n", key)
	q, e := query(key)
	if e != nil {
		log.Printf("queryHandler: query[%q] failed: %s\n", key, e)
		return
	}
	if e := queryTem.Execute(res, q); e != nil {
		log.Printf("queryHandler: %s\n", e)
	}
}

func StartSever() error {
	return http.ListenAndServe(":8080", r)
}
