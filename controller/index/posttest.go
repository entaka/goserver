package index

import (
	"encoding/json"
	"fmt"
	"goserver/system/db"
	"log"
	"net/http"
	"strings"
)

type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

func init() {

	http.HandleFunc("/posttest/", PostTest)
}

func PostTest(w http.ResponseWriter, r *http.Request) {
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, _ := json.Marshal(group)

	r.ParseForm()
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	Update()

	fmt.Fprintf(w, string(b))
	fmt.Println(string(b))
}

type Counting struct {
	Id  int
	Sum int
}

func Update() {
	orm := db.Beedb()

	var counting Counting
	//Whereは2つの引数を受け取ります。int型の引数をサポートします。
	orm.Where(1).Find(&counting)
	log.Printf("%#v", counting)

	t := make(map[string]interface{})
	t["sum"] = counting.Sum + 1
	orm.SetTable("counting").SetPK("id").Where(1).Update(t)

}
