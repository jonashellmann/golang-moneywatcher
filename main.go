package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()

	staticFileDirectory := http.Dir("./a/")
	staticFileHandler := http.StripPrefix("/a/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/a/").Handler(staticFileHandler).Methods("GET")

	r.HandleFunc("/region", getRegionHandler).Methods("GET")
	r.HandleFunc("/category", getCategoryHandler).Methods("GET")
	r.HandleFunc("/recipient", getRecipientHandler).Methods("GET")
	r.HandleFunc("/expense", getExpenseHandler).Methods("GET")
	return r
}

func main() {
	createDatabase()

	r := newRouter()
	http.ListenAndServe(":5656", r)
}

func createDatabase() {
	configuration, err1 := ReadConfiguration()

	if err1 != nil {
	        panic(err1)
	}

	connString := configuration.User + ":" + configuration.Password + "@/" + configuration.Database + "?parseTime=true"
	db, err := sql.Open("mysql", connString)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	InitStore(&dbStore{db: db})
	err = store.CreateStorage()

	if err != nil {
		panic(err)
	}
}
