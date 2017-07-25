package database

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	// "encoding/json"
	"goduan/model"
	"log"
	"time"
)

const (
	MongoDBHosts = "bk0uk7n6pqkfykg-mongodb.services.clever-cloud.com"
	AuthDatabase = "bk0uk7n6pqkfykg"
	AuthUserName = "uon3frjggc9avs6"
	AuthPassword = "AYDl8wNSbDYtEXGPLdFr"
	TestDatabase = "goinggo"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController() *UserController {
	mogoDbInfo := &mgo.DialInfo{
		Addrs:    []string{MongoDBHosts},
		Timeout:  60 * time.Second,
		Database: AuthDatabase,
		Username: AuthUserName,
		Password: AuthPassword,
	}

	session, err := mgo.DialWithInfo(mogoDbInfo)
	if err != nil {
		log.Fatalf("CreateSession: %s\n", err)
	} else {
		session.SetMode(mgo.Monotonic, true)
	}

	return &UserController{session}
}

func (uc *UserController) InsertAnUser(user model.User) (bool,string){
	exist:=uc.session.DB(AuthDatabase).C("users").Find(bson.M{"email":user.Email}).Select(nil).One(&model.User{})
	if exist==nil{
		return false,"Email đã được sử dụng"
	}
	err:=uc.session.DB(AuthDatabase).C("users").Insert(user)
	if err != nil {
		return false,"False"
	}
	return true,"Success"
}

func (uc *UserController) GetAnUser(email, pass string) (bool,model.User) {
	c := uc.session.DB(AuthDatabase).C("users")
	user := model.User{}
	err := c.Find(bson.M{"email": email}).Select(nil).One(&user)
	if err != nil{
		return false,user
	}

	return true,user
}

func (uc *UserController) InsertACategory(category model.Category) bool{
	err:=uc.session.DB(AuthDatabase).C("category").Insert(category)
		if err != nil {
		return false
	}

	return true
}

func (uc *UserController) InsertAProfile(profile model.Profile) bool{
	err:=uc.session.DB(AuthDatabase).C("profile").Insert(profile)
		if err != nil {
		return false
	}

	return true
}

func (uc *UserController) GetProfileWithUser(profile_id string) (bool,model.Profile) {
	c := uc.session.DB(AuthDatabase).C("profile")
	pro := model.Profile{}
	err := c.Find(bson.M{"profile_id": profile_id}).Select(nil).One(&pro)
	if err != nil{
		return false,pro
	}

	return true,pro
}


func(uc *UserController) GetAllCategory()	([]model.Category,int){
	 var results []model.Category	
	c:=uc.session.DB(AuthDatabase).C("category")
	err:=c.Find(nil).All(&results)
	if err!=nil {
		return results,0
	}
	return results,1
}

func (uc *UserController) GetUsersCount() int {
	count,_:=uc.session.DB(AuthDatabase).C("users").Count()

	return count
}

func (uc *UserController) GetCategoryCount() int {
	count,_:=uc.session.DB(AuthDatabase).C("category").Count()

	return count
}

func (uc *UserController) UpdateProfileValue(pro Profile) (bool,mess){
	
}