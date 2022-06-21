package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alireza-ghavabesh/sql-server-test/connection"
)

func StartServer() {

	http.HandleFunc("/create_table", func(w http.ResponseWriter, req *http.Request) {
		query := req.URL.Query()
		tableName := query["name"][0]
		result := connection.CreateTable(tableName)
		if result == tableName {
			fmt.Fprintf(w, "table %s created.", tableName)
		} else {
			fmt.Fprintf(w, "there is a problem for name <(%s)>...!", tableName)
		}
	})

	http.HandleFunc("/drop_table", func(w http.ResponseWriter, req *http.Request) {
		query := req.URL.Query()
		tableName := query["name"][0]
		result := connection.DropTable(tableName)
		if result == tableName {
			fmt.Fprintf(w, "table %s droped.", tableName)
		} else {
			fmt.Fprintf(w, "result for <<%s>>: %s", tableName, result)
		}
	})

	http.HandleFunc("/insert", func(w http.ResponseWriter, req *http.Request) {
		query := req.URL.Query()
		tableName := query["name"][0]
		result := connection.Insert(tableName)
		if result == tableName {
			fmt.Fprintf(w, "inserted")
		} else {
			fmt.Fprintf(w, "result for <<%s>>: %s", tableName, result)
		}
	})

	http.HandleFunc("/select", func(w http.ResponseWriter, req *http.Request) {
		query := req.URL.Query()
		tableName := query["name"][0]
		result := connection.Select(tableName)
		if result == tableName {
			fmt.Fprintf(w, "selected")
		} else {
			fmt.Fprintf(w, "result for <<%s>>: %s", tableName, result)
		}
	})

	fmt.Println("server started: http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
