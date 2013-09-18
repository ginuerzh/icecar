// Message
package models

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type Message struct {
	Id           bson.ObjectId `bson:"_id"`
	Receiver     string
	Sender       string
	Message_text string
	Unread       bool
	Post_time    time.Time
}
