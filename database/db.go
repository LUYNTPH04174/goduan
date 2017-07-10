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

type UserController struct{
	session *mgo.Session
}

func NewUserController() *UserController{
	mogoDbInfo:=&mgo.DialInfo{
	Addrs:    []string{MongoDBHosts},
			Timeout:  60 * time.Second,
			Database: AuthDatabase,
			Username: AuthUserName,
			Password: AuthPassword,
	}

	session,err:=mgo.DialWithInfo(mogoDbInfo);
	if err != nil {
			log.Fatalf("CreateSession: %s\n", err)
	}else{
		session.SetMode(mgo.Monotonic,true)
	}
	
	return &UserController{session}
}

func (uc *UserController)InsertAnUser(user model.User) {
	uc.session.DB(AuthDatabase).C("users").Insert(user)

}

func (uc *UserController) GetAnUser(email,pass string) bool{
	c:=uc.session.DB(AuthDatabase).C("users")
	result:=model.User{}

	err := c.Find(bson.M{"sdt": email}).Select(bson.M{"password": pass}).One(&result)
	if err!=nil{
		return false
	}

	return true
}