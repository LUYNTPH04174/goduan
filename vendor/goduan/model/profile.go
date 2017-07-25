package model

import (
 "gopkg.in/mgo.v2/bson"
)

type Profile struct {
	Id  bson.ObjectId `json:"id" bson:"_id"`
	Profile_id   string `json:"profile_id" bson:"profile_id"`
	Name string `form:"name" json:"name" bson:"name"`
	Phone_number 	string `form:"phone_number" json:"phone_number" bson:"phone_number"`
	Contact_emal   	string  `form:"contact_emal" json:"contact_emal" bson:"contact_emal"`
	Address			string	`form:"address" json:"address" bson:"address"`
	Create_at		string 	`form:"create_at" json:"create_at" bson:"create_at"`
	Update_at 		string 	`form:"update_at" json:"update_at" bson:"update_at"`
	Avatar_id		string 	`form:"avatar_id" json:"avatar_id" bson:"avatar_id"`
}

func (profile *Profile) SetValueProfile(id,name,phone_number,contact_emal,address,create_at,update_at string) {
	profile.Id= bson.NewObjectId()
	profile.Profile_id =id
	profile.Name=name
	profile.Phone_number=phone_number
	profile.Contact_emal = contact_emal
	profile.Address = address
	profile.Create_at = create_at
	profile.Update_at = update_at
}

