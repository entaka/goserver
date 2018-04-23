package main

import (
	"fmt"
	"goserver/controller"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //オプションを解析します。デフォルトでは解析しません。
	fmt.Println(r.Form) //このデータはサーバのプリント情報に出力されます。
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "AAAHello astaxie!<button>aa</button>") //ここでwに入るものがクライアントに出力されます。
}

func testFunc(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)

	fmt.Fprintf(w, "go go test")
}

func main() {

	http.HandleFunc("/", sayhelloName)  //アクセスのルーティングを設定します。
	http.HandleFunc("/test/", testFunc) // テスト用
	controller.Controller()

	// http://localhost:9090/js/test.js
	//	http.Handle("/js/", http.FileServer(http.Dir("../../templates/goserver/")))
	http.Handle("/js/", http.FileServer(http.Dir("assets/")))

	err := http.ListenAndServe(":9090", nil) //監視するポートを設定します。
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
