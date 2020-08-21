package main

import (
	"database/sql"
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)

func openDB() *sql.DB{
	db,err := sql.Open("mysql","root:123456@tcp(121.41.73.98:3306)/shorturl")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

var db = openDB()

func main() {
	http.HandleFunc("/",getShortURL)
	err := http.ListenAndServe(":9091",nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

