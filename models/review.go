// review.go
package models

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type Review struct {
	Id           bson.ObjectId `bson:"_id"`
	Receiver     string
	Sender       string
	Message_text string
	Unread       bool
	Carid        string
	Post_time    time.Time
}
