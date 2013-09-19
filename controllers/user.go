// user
package controllers

import (
	"encoding/json"
	//"fmt"
	"github.com/ginuerzh/icecar/models"
	//"labix.org/v2/mgo"
	//"labix.org/v2/mgo/bson"
	"github.com/ginuerzh/icecar/errors"
	"time"
)

type UserController struct {
	BaseController
}

func (this *UserController) Register() {
	var user models.User

	//json.Unmarshal(this.Ctx.RequestBody, &user)
	decoder := json.NewDecoder(this.Ctx.Request.Body)
	if err := decoder.Decode(&user); err != nil {
		this.Data["json"] = this.response(nil, &errors.JsonError)
		this.ServeJson()
		return
	}

	user.Userid = user.Email
	user.Password = this.md5(user.Password)
	user.RegTime = time.Now()

	if err := user.Save(); err != nil {
		this.Data["json"] = this.response(nil, err)
	} else {
		user.AccessToken = this.uuid()
		r := map[string]string{"access_token": user.AccessToken}
		this.Data["json"] = this.response(r, nil)
	}

	user.SetOnline(true)

	this.ServeJson()
}

func (this *UserController) Login() {
	var user models.User

	//json.Unmarshal(this.Ctx.RequestBody, &user)
	decoder := json.NewDecoder(this.Ctx.Request.Body)
	if err := decoder.Decode(&user); err != nil {
		this.Data["json"] = this.response(nil, &errors.JsonError)
		this.ServeJson()
		return
	}

	user.Password = this.md5(user.Password)

	if err := user.LoginCheck(); err != nil {
		this.Data["json"] = this.response(nil, err)
		this.ServeJson()
		return
	}

	user.AccessToken = this.uuid()
	r := map[string]string{"access_token": user.AccessToken}
	this.Data["json"] = this.response(r, nil)

	user.SetOnline(true)

	this.ServeJson()
}

func (this *UserController) Logout() {
	this.Ctx.WriteString("logout")
}

func (this *UserController) UserInfo() {
	var user models.User

	//json.Unmarshal(this.Ctx.RequestBody, &user)
	decoder := json.NewDecoder(this.Ctx.Request.Body)
	if err := decoder.Decode(&user); err != nil {
		this.Data["json"] = this.response(nil, &errors.JsonError)
		this.ServeJson()
		return
	}

	if err := user.Load(); err != nil {
		this.Data["json"] = this.response(nil, err)
		this.ServeJson()
		return
	}

	resp := make(map[string]interface{})

	resp["userid"] = user.Userid
	resp["nikename"] = user.Nickname
	resp["phone_number"] = user.Phone
	resp["profile_image"] = user.Profile
	resp["photos"] = user.Photos
	resp["register_time"] = user.RegTime.Format("2006-01-02 15:04:05")
	/*
		c = DB.C("CarInfo")
		var cars []string
		if err := c.Find(bson.M{"host_userid": user.Userid}).Distinct("_id", &cars); !err.id {
			this.Data["json"] = this.errorResp(DbError)
			this.ServeJson()
			return
		}
		resp["car_ids"] = cars

		c = DB.C("Reviews")
		var review models.Review
		query := c.Find(bson.M{"receiver": user.Userid}).Sort("-post_time")
		resp["review_count"], _ = query.Count()
		resp["picked_review"] = nil
		resp["picked_review_author"] = nil
		if err := query.One(&review); err == nil {
			resp["picked_review"] = review.Message_text
			resp["picked_review_author"] = review.Sender
			resp["picked_review_car"] = review.Carid
		}
	*/
	this.Data["json"] = this.response(resp, nil)
	this.ServeJson()
}

func (this *UserController) NewsCount() {
	//c := DB.C("Messages")
}
