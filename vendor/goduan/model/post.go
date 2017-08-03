package model

import (
 "gopkg.in/mgo.v2/bson"
 //"strings"
)

type PostDetail struct {
	Id  bson.ObjectId `json:"id" bson:"_id"`
	Title string `form:"title" json:"title" bson:"title"`
	Description	string	`form:"description" json:"description" bson:"description"`
	Category_Id	string 	`form:"category_id" json:"category_id" bson:"category_id"`
	Created_at	string 	`form:"created_at" json:"created_at" bson:"created_at"`
	Created_by 	string 	`form:"created_by" json:"created_by" bson:"created_by"`
	Profile_Id	string 	`form:"profile_id" json:"profile_id" bson:"profile_id"`
	Images 	string `json:"images" bson:"images"`
	Place string 	`form:"place" json:"place" bson:"place"`
	Requirement	string `form:"requirement" json:"requirement" bson:"requirement"`
	Benefits	string `form:"benefits" json:"benefits" bson:"benefits"`
	Time_limited string `form:"time_limited" json:"time_limited" bson:"time_limited"`
	Contact string `form:"contact" json:"contact" bson:"contact"`
}

func (detail *PostDetail) SetValueInfo(title,category_id,create_at,create_by,profile_id string) {
	detail.Id=bson.NewObjectId()
	detail.Title=title
	detail.Category_Id=category_id
	detail.Created_at=create_at
	detail.Created_by=create_by
	detail.Profile_Id=profile_id
}


func (detail *PostDetail) SetValueDetail(place,description,requirement,benefits,time_limited,contact string){
	detail.Place=place
	detail.Description=description
	detail.Requirement=requirement
	detail.Benefits=benefits
	detail.Time_limited=time_limited
	detail.Contact=contact
}

func (detail *PostDetail) SetImages(indata string){
	detail.Images=indata
	// detail.Images=strings.Split(indata,",")
}