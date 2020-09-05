package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func getHTML(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("shortURL.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, nil)
}

func getShortURL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")

	err := r.ParseForm()
	if err != nil {
		log.Fatal("request parse:", err)
	}
	longurl := r.Form["longurl"][0]
	fmt.Println(longurl)
	//查找缓存
	cacheRes := Cache.find(longurl)
	if len(cacheRes) > 0 {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "http://"+ServerIP+":9091/"+cacheRes)
		return
	}
	//缓存未命中
	stmt, err := db.Prepare("insert into record(longurl,times) value(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(longurl, 0)
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	//id转短网址
	shortUrl := id2url(lastId)
	_, err = db.Exec("update record set shorturl = '"+shortUrl+"' where id = ?", lastId)
	if err != nil {
		log.Fatal(err)
	}
	Cache.add(longurl, shortUrl)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "http://"+ServerIP+":9091/"+shortUrl)
}

func getLongURL(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal("parseerr:", err)
	}
	shortUrl := r.URL.Path
	shortUrl = shortUrl[1:]
	ip := strings.Split(r.RemoteAddr, ":")[0]
	//筛选短网址，ip
	if len(shortUrl) > 6 || len(shortUrl) == 0 || checkIP(ip) {
		http.Redirect(w, r, "http://"+ServerIP+":9091/shortURL", http.StatusFound)
		return
	}

	id := url2id(shortUrl)
	var longUrl string
	err = db.QueryRow("select longurl from record where id=?", id).Scan(&longUrl)
	if err != nil {
		http.Redirect(w, r, "http://"+ServerIP+":9091/shortURL", http.StatusFound)
		return
	}
	http.Redirect(w, r, longUrl, http.StatusFound)
	_, err = db.Exec("update record set times = times + 1 where id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
}
