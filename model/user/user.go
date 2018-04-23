package user

import (
	"goserver/system/db"
	"time"
)

// User Table struct
type User struct {
	ID      int
	Name    string
	Detail  string
	Created time.Time
}

// Insert Method
func Insert(arg User) int64 {
	orm := db.Beedb()

	var saveone User
	saveone.Name = arg.Name
	saveone.Detail = arg.Detail
	saveone.Created = time.Now()
	orm.Save(&saveone)

	return int64(saveone.ID)
}
