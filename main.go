package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func openDB() *sql.DB {
	db, err := sql.Open("mysql", "user:password@tcp(host:port)/shorturl")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

var db = openDB()
var ServerIP string = "127.0.0.1"
var Cache LRUcache

func main() {
	Cache.init(62 * 62 * 62)
	go ipControllor()

	http.HandleFunc("/shortURL", getHTML)
	http.HandleFunc("/getShortURL", getShortURL)
	http.HandleFunc("/", getLongURL)
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
