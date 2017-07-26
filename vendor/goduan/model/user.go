package model

import (
	// "encoding/json"
 "gopkg.in/mgo.v2/bson"
)

type User struct {
	Id	bson.ObjectId `json:"id" bson:"_id"`
	Id_user string `json:"user_id" bson:"user_id"`
	Email string `form:"email" json:"email" bson:"email"`
	Password string `form:"password" json:"password" bson:"password"`
	Facebook_id string `form:"facebook_id" json:"facebook_id" bson:"facebook_id"`
	Profile_id string `form:"profile_id" json:"profile_id" bson:"profile_id"`
	Create_at string `form:"create_at" json:"create_at" bson:"create_at"`
	Token string `json:"token" bson:"token"`
}

func (user *User) SetValueUser(id,email,password,facebook_id,profile_id,create_at string) {
	user.Id =bson.NewObjectId()
	user.Id_user = id
	user.Email = email
	user.Password = password
	user.Facebook_id = facebook_id
	user.Profile_id = profile_id
	user.Create_at = create_at
	user.Token=user.Id.Hex()
}

func (user *User) GetEmail() string {
	return user.Email
}

func (user *User) GetPassWord() string {
	return user.Password
}

func (user *User) GetFacebook() string {
	return user.Facebook_id
}
