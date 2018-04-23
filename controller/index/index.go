package index

import (
	"fmt"
	"goserver/model/user_status"
	"html/template"
	"log"
	"net/http"
)

func Test1Func(w http.ResponseWriter, r *http.Request) {
	// テンプレートをパース
	t := template.Must(template.ParseFiles("view/template01.html.tpl"))

	str := "Sample Message"

	fmt.Fprintf(w, "go go test")

	// テンプレートを描画
	if err := t.ExecuteTemplate(w, "template01.html.tpl", str); err != nil {
		log.Fatal(err)
	}
}

func Test2Func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Test2Func")
	// テンプレートをパース
	t := template.Must(template.ParseFiles("view/template02.tpl.html"))

	m := map[string]int{
		"key1": 101,
		"key2": 202,
		"key3": 303,
		"key4": -404,
	}

	// テンプレートを描画
	if err := t.ExecuteTemplate(w, "template02.tpl.html", m); err != nil {
		log.Fatal(err)
	}
}

func Test3Func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Test3Func")
	// テンプレートをパース
	//t := template.Must(template.ParseFiles("../../templates/goserver/template03.tpl.html"))
	t := template.Must(template.ParseFiles("view/template03.tpl.html"))

	m := map[string]int{
		"key1": 101,
		"key2": 202,
		"key3": 303,
		"key4": -404,
	}
	fmt.Println("INSERT START")
	user_status.Insert(user_status.User_status{
		User_id:     1,
		Status_name: "SPEED",
		Value:       "255",
	})
	fmt.Println("INSERT END")

	// テンプレートを描画
	if err := t.ExecuteTemplate(w, "template03.tpl.html", m); err != nil {
		log.Fatal(err)
	}
}

func Test4Func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Test4Func")
	// テンプレートをパース
	t := template.Must(template.ParseFiles("view/template02.tpl.html"))

	fmt.Println("INSERT START")
	user_status.Insert(user_status.User_status{
		User_id:     1,
		Status_name: "SPEED",
		Value:       "255",
	})
	fmt.Println("INSERT END")

	m := map[string]int{
		"key1": 101,
		"key2": 202,
		"key3": 303,
		"key4": -404,
	}

	// テンプレートを描画
	if err := t.ExecuteTemplate(w, "template02.tpl.html", m); err != nil {
		log.Fatal(err)
	}
}
