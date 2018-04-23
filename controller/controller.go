package controller

import (
	"fmt"
	"goserver/controller/index"
	"net/http"
)

func Controller() {
	fmt.Println("Controller S")
	http.HandleFunc("/test1/", index.Test1Func) // テスト用
	http.HandleFunc("/test2/", index.Test2Func)
	http.HandleFunc("/test3/", index.Test3Func)
	http.HandleFunc("/test4/", index.Test4Func)

	fmt.Println("Controller E")
}
