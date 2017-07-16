package model

import (
	// "encoding/json"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id         bson.ObjectId `json:"id" bson:"_id"`
	Sdt        string        `form:"sdt" json:"sdt" bson:"sdt"`
	Password   string        `form:"password" json:"password" bson:"password"`
	Facebookid string        `form:"facebookid" json:"facebookid" bson:"facebookid"`
}

func (user *User) SetValue(a, b, c string) {
	user.Id = bson.NewObjectId()
	user.Sdt = a
	user.Password = b
	user.Facebookid = c
}

func (user *User) GetSDT() string {
	return user.Sdt
}

func (user *User) GetPassWord() string {
	return user.Password
}

func (user *User) GetFacebook() string {
	return user.Facebookid
}
