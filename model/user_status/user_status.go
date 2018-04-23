package user_status

import (
	"fmt"
	"goserver/system/db"
	"time"
)

type User_status struct {
	Id          int
	User_id     int
	Status_name string
	Value       string
	Created     time.Time
}

func Insert(arg User_status) {
	orm := db.Beedb()

	var saveone User_status
	saveone.User_id = arg.User_id
	saveone.Status_name = arg.Status_name
	saveone.Value = arg.Value
	saveone.Created = time.Now()
	orm.Save(&saveone)
	fmt.Println("ORM SAVE")
	orm.Db.Close()

}
