package main

import (
	"fmt"
	"log"
	"net/http"
)

func getShortURL(w http.ResponseWriter,r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")

	err := r.ParseForm()
	if err !=  nil {
		log.Fatal("request parse:",err)
	}
	longurl := r.Form["longurl"][0]
	fmt.Println(longurl)
	//检查缓存

	//缓存未命中
	stmt,err := db.Prepare("insert into record(longurl,times) value(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	res,err := stmt.Exec(longurl,0)
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lastId)
	//id转短网址

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w,"getlongurl")
}
