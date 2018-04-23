package db

import (
	"database/sql"

	"github.com/astaxie/beedb"
	_ "github.com/go-sql-driver/mysql"
)

func Beedb() (orm beedb.Model) {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/gotest?charset=utf8")
	if err != nil {
		panic(err)
	}

	orm = beedb.New(db)
	return orm
}

/** 以下未使用 **/

//func Insert(sql string, args ...interface{}) int64 {
//	//データの挿入
//	stmt, err := db.Prepare(sql)
//	CheckErr(err)

//	res, err := stmt.Exec(args...)
//	CheckErr(err)

//	id, err := res.LastInsertId()
//	CheckErr(err)

//	return id
//}

//func Update(id int64) int64 {
//	stmt, err := db.Prepare("update userinfo set username=? where uid=?")
//	CheckErr(err)

//	res, err := stmt.Exec("astaxieupdate", id)
//	CheckErr(err)

//	affect, err := res.RowsAffected()
//	CheckErr(err)

//	return affect
//}

//func Get() (rows *sql.Rows) {
//	rows, err := db.Query("SELECT * FROM userinfo")
//	CheckErr(err)
//	return rows
//}

//func Delete(id int64) int64 {
//	stmt, err := db.Prepare("delete from userinfo where uid=?")
//	CheckErr(err)

//	res, err := stmt.Exec(id)
//	CheckErr(err)

//	affect, err := res.RowsAffected()
//	CheckErr(err)

//	return affect
//}

//func CheckErr(err error) {
//	if err != nil {
//		panic(err)
//	}
//}
