package model

import (
 "gopkg.in/mgo.v2/bson"
)

type Profile struct {
	Id  bson.ObjectId `json:"id" bson:"_id"`
	Profile_Stt   string `json:"profile_stt" bson:"profile_stt"`
	Name string `form:"name" json:"name" bson:"name"`
	Phone_number 	string `form:"phone_number" json:"phone_number" bson:"phone_number"`
	Contact_email   	string  `form:"contact_email" json:"contact_email" bson:"contact_email"`
	Address			string	`form:"address" json:"address" bson:"address"`
	Create_at		string 	`form:"create_at" json:"create_at" bson:"create_at"`
	Updated_at 		string 	`form:"updated_at" json:"updated_at" bson:"updated_at"`
	Avatar_id		string 	`form:"avatar_id" json:"avatar_id" bson:"avatar_id"`
}

func (profile *Profile) SetValueProfile(id,name,phone_number,contact_email,address,create_at,update_at string) {
	profile.Id= bson.NewObjectId()
	profile.Profile_Stt =id
	profile.Name=name
	profile.Phone_number=phone_number
	profile.Contact_email = contact_email
	profile.Address = address
	profile.Create_at = create_at
	profile.Updated_at = update_at
}

func (profile *Profile) SetAvatarID(linkava string) {
	profile.Avatar_id=linkava
}