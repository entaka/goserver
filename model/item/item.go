package item

import (
	"database/sql"
	//"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Argments struct {
	Id      int
	Name    string
	Detail  string
	Created string
}

var db *sql.DB

func init() {
	db = Access()
}

func Access() (db *sql.DB) {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/gotest?charset=utf8")
	if err != nil {
		panic(err)
	}
	return db
}

func Insert(arg Argments) int64 {

	//データの挿入

	sql := `
      INSERT item 
      SET
      name=?,
      detail=?,
      created=?
    `
	now := time.Now().Format("2006-01-02 15:04:05")

	//データの挿入
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(arg.Name, arg.Detail, now)
	if err != nil {
		panic(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}

	return id
}
