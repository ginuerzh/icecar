// user
package models

import (
	"icecar/errors"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

type AuthType string
type AuthStatus string

const (
	AuthPhone          = AuthType("phone")
	AuthIdCard         = AuthType("id_card")
	AuthDriveLicense   = AuthType("driver_license")
	AuthVehicleLicense = AuthType("vehicle_license")
	AuthOthers         = AuthType("vehicle_license")

	AuthVerified   = AuthStatus("verified")
	AuthVerifying  = AuthStatus("verifying")
	AuthUnverified = AuthStatus("unverified")
	AuthRefused    = AuthStatus("refused")
)

type User struct {
	Id          bson.ObjectId `bson:"_id"`
	Userid      string
	Password    string
	Email       string
	Nickname    string            `json:"nikename"`
	Phone       string            `bson:"phone,omitempty"`
	Profile     string            `bson:",omitempty"`
	Photos      []string          `bson:",omitempty"`
	RegTime     time.Time         `bson:"reg_time"`
	AuthStatus  map[string]string `bson:"auth_status,omitempty"`
	Role        string            `bson:"role,omitempty"`
	Online      bool
	LastAccess  time.Time `bson:"last_access"`
	AccessToken string    `bson:"access_token,omitempty"`
}

type UserAuth struct {
	Images []string `bson:"auth_images"`
	Info   string   `bson:"auth_info"`
	Desc   string   `bson:"auth_description"`
	Type   AuthType
	Status AuthStatus
}

func (this *User) Load() (e *errors.Error) {
	c := DB.C(C_User)

	if err := c.Find(bson.M{"userid": this.Userid}).One(this); err != nil {
		if err == mgo.ErrNotFound {
			e = &errors.UserNotFoundError
		} else {
			e = &errors.DbError
		}
	}
	return
}

func (this *User) Exists() (b bool) {
	c := DB.C(C_User)

	count, _ := c.Find(bson.M{"userid": this.Userid}).Count()

	if count > 0 {
		b = true
	}

	return
}

func (this *User) Save() (e *errors.Error) {
	c := DB.C(C_User)

	if this.Exists() {
		e = &errors.UserExistError
		return
	}

	this.Id = bson.NewObjectId()
	if err := c.Insert(this); err != nil {
		e = &errors.DbError
	}

	return
}

func (this *User) LoginCheck() (e *errors.Error) {
	c := DB.C(C_User)

	if err := c.Find(bson.M{"userid": this.Userid, "password": this.Password}).One(this); err != nil {
		if err == mgo.ErrNotFound {
			e = &errors.UserNotFoundError
		} else {
			e = &errors.DbError
		}
	}
	return
}

func (this *User) SetOnline(online bool) (e *errors.Error) {
	c := DB.C(C_User)

	if online && len(this.AccessToken) == 0 {
		e = &errors.AccessError
		return
	}

	change := bson.M{
		"$set": bson.M{
			"online":       online,
			"last_access":  bson.Now(),
			"access_token": this.AccessToken,
		},
	}

	if err := c.UpdateId(this.Id, change); err != nil {
		e = &errors.DbError
		return
	}

	this.Online = online
	return
}
