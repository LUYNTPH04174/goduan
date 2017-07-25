package model

import (
	// "encoding/json"
 "gopkg.in/mgo.v2/bson"
)

type User struct {
	Id	bson.ObjectId `json:"id" bson:"_id"`
	Id_user string `json:"id" bson:"id"`
	Email string `form:"email" json:"email" bson:"email"`
	Password string `form:"password" json:"password" bson:"password"`
	Facebook_id string `form:"facebook_id" json:"facebook_id" bson:"facebook_id"`
	Profile_id string `form:"profile_id" json:"profile_id" bson:"profile_id"`
	Create_at string `form:"create_at" json:"create_at" bson:"create_at"`
	Update_at string `form:"update_at" json:"update_at" bson:"update_at"`
	Create_by string `form:"create_by" json:"create_by" bson:"create_by"`
	Update_by string `form:"update_by" json:"update_by" bson:"update_by"`
	Profile Profile `json:"profile" bson:"profile"`
	
}

func (user *User) SetProf(profile Profile) {
	user.Profile=profile
}
func (user *User) SetValueUser(id,email,password,facebook_id,profile_id,create_at,update_at,create_by,update_by string) {
	user.Id =bson.NewObjectId()
	user.Id_user = id
	user.Email = email
	user.Password = password
	user.Facebook_id = facebook_id
	user.Profile_id = profile_id
	user.Create_at = create_at
	user.Update_at = update_at
	user.Create_by = create_by
	user.Update_by = update_by
}

func (user *User) GetSDT() string {
	return user.Email
}

func (user *User) GetPassWord() string {
	return user.Password
}

func (user *User) GetFacebook() string {
	return user.Facebook_id
}
