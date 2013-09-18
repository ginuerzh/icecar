// model.go
package models

import (
	"labix.org/v2/mgo"
)

var (
	DB       *mgo.Database
	C_User   = "UserInfo"
	C_Review = "Reviews"
)
