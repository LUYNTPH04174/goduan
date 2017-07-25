package model
import (
	"gopkg.in/mgo.v2/bson"
)
type Category struct{
 	Id  bson.ObjectId `json:"id" bson:"_id"`
	Id_category string `form:"id_category" json:"id_category" bson:"id_category"`
	Category_name string `form"category_name" json:"category_name" bson:"category_name"`
	Create_at string `form:"create_at" json:"create_at" bson:"create_at"`
	Updated_at string `form:"updated_at" json:"updated_at" bson:"updated_at"`
	Create_by		string 		`form:"create_by" 		json:"create_by" 	bson:"create_by"`
	Update_by		string 		`form:"update_by" 		json:"update_by" 	bson:"update_by"`
	Icon_id string `form:"icon_id" json:"icon_id" bson:"icon_id"`
}	

func (category *Category) SetValueCategory(id, name, create_at,updated_at,icon_id string) {
	category.Id = bson.NewObjectId()
	category.Id_category = id
	category.Category_name = name
	category.Create_at = create_at
	category.Updated_at = updated_at
	category.Icon_id = icon_id
}

func (category *Category) GetID() string {
	return category.Id_category
}

func (category *Category) GetCategoryName() string {
	return category.Category_name
}

func (category *Category) GetCreateTime() string {
	return category.Create_at
}

func (category *Category) GetUpdateTime() string {
	return category.Updated_at
}

func (category *Category) GetIconId() string {
	return category.Icon_id
}
