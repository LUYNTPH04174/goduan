package model

import (
 "gopkg.in/mgo.v2/bson"
)

type PostDetail struct {
	Id  bson.ObjectId `json:"id" bson:"_id"`
	Title string `form:"title" json:"title" bson:"title"`
	Description	string	`form:"address" json:"address" bson:"address"`
	Category_Id	string 	`json:"category_id" bson:"category_id"`
	Created_at	string 	`json:"create_at" bson:"create_at"`
	Created_by 	string 	`json:"create_by" bson:"create_by"`
	Profile_Id	string 	`json:"profile_id" bson:"profile_id"`
	Images 	[]string `json:"images" bson:"images"`
}

func (detail *PostDetail) SetValueDetail(title,des,category_id,create_at,create_by,profile_id string) {
	detail.Id=bson.NewObjectId()
	detail.Title=title
	detail.Description=des
	detail.Category_Id=category_id
	detail.Created_at=create_at
	detail.Created_by=create_by
	detail.Profile_Id=profile_id
}

