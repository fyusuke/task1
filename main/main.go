package main

import (
	"log"
	"net/http"
	"html/template"
	"time"
	"strings"
	"fmt"
)

// 時刻を日本時間に合わせる
func init() {
	const location = "Asia/Tokyo"
    loc, err := time.LoadLocation(location)
    if err != nil {
        loc = time.FixedZone(location, 9*60*60)
    }
    time.Local = loc
}

func main() {

	port := "8080"

	http.HandleFunc("/", mainHandler)
	log.Printf("Server listening on http://localhost:%s/", port)
	log.Print(http.ListenAndServe(":"+port, nil))
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	const ip string = "172.16.238.3:"
	tpl := template.Must(template.ParseFiles("template/index.html"))

	// 朝７時から夜７時の間の広告
	Banner := "http://placehold.jp/150x50.png" 

	nowHour := time.Now().Hour()
	isSpecificIP := strings.Contains(getIP(r), ip)
	if(isSpecificIP==true) { fmt.Println("特定のIPアドレスからのアクセスが来ました") }
	// 特定のIPアドレスからのアクセス or 夜７時から朝７時の間の広告
	if isSpecificIP || nowHour < 7 || nowHour >= 19 {
		Banner = "http://placehold.jp/300x150.png"
	}

	// fmt.Println("nowBanner: "+Banner)
	// fmt.Println("getIP(r): "+getIP(r))
	
	err := tpl.Execute(w, Banner)
	if  err != nil {
		panic(err.Error())
	}
}

// IPアドレスの取得
func getIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}