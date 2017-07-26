package route

import (
	"github.com/gin-gonic/gin"
	"goduan/database"
	"goduan/model"
	"net/http"
	"strconv"
	// "encoding/json"
	"time"
	"fmt"
)

var controller = database.NewUserController()

func HomeDefault(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"status":  http.StatusOK})
}


func InsertAnUserRouter(c *gin.Context) {
	email := c.PostForm("email")
	pass := c.PostForm("pass")
	facebookid := c.PostForm("facebook_id")
	user_id:=strconv.Itoa(controller.GetUsersCount()+1)
	create_at:=time.Now().Format(time.RFC1123)
	name:=c.PostForm("name")
	user := model.User{}
	// SetValueProfile(id,name,phone_number,contact_emal,address,create_at,update_at)
	profile:=model.Profile{}
	profile.SetValueProfile(user_id,name,"","","",create_at,"")
	profile_id:=profile.Id.Hex()
	//SetValueUser(email,password,facebook_id,profile_id,create_at string)
	user.SetValueUser(user_id,email, pass, facebookid,profile_id,create_at)

	err,message:=controller.InsertAnUser(user)
	if err&&controller.InsertAProfile(profile){
		c.JSON(http.StatusOK, gin.H{
		"message":message,
		"status":  http.StatusOK})
	}else{
		c.JSON(http.StatusOK,gin.H{
		"message":message,
		"status":http.StatusNotFound})
	}

}

func GetLoginAnUser(c *gin.Context) {
	email := c.Query("email")
	password := c.Query("pass")
	if email==""||password==""{
		facebookid:=c.Query("facebook_id")
		succ,user:=controller.GetAnUserByFace(facebookid)
		succ2,pro:=controller.GetProfileWithUser(user.Id_user)
		fmt.Println(succ,succ2,user.Id_user)
		if succ &&succ2{
			c.JSON(http.StatusOK,gin.H{
				"user":user,
				"profile":pro})
		}else{
			c.JSON(http.StatusNotFound,gin.H{"message":"Thông tin chưa chính xác",
				"status":http.StatusOK})
		}
	}else{
		succ,user:=controller.GetAnUser(email, password)
		succ2,pro:=controller.GetProfileWithUser(user.Id_user)
		fmt.Println(succ,succ2,user.Id_user)
		if succ &&succ2{
			c.JSON(http.StatusOK,gin.H{
				"user":user,
				"profile":pro})
		}else{
			c.JSON(http.StatusNotFound,gin.H{"message":"Thông tin chưa chính xác",
				"status":http.StatusOK})
		}
	}
}

func InsertACategory(c *gin.Context) {
	id:=strconv.Itoa(controller.GetCategoryCount()+1)
	name:=c.PostForm("category_name")
	create_at:=c.PostForm("create_at")
	updated_at:=c.PostForm("updated_at")
	icon_id:=c.PostForm("icon_id")

	category:=model.Category{}
	category.SetValueCategory(id,name,create_at,updated_at,icon_id)
	if controller.InsertACategory(category){
		c.JSON(http.StatusOK,gin.H{
		"message":"Success",
		"status":http.StatusOK})
	}else{
		c.JSON(http.StatusOK,gin.H{
		"message":"False",
		"status":http.StatusOK})
	}
}

func GetCategorys(c *gin.Context) {
	results,err:=controller.GetAllCategory()
	if err==1{
		c.JSON(http.StatusOK,results)
	}else{
		c.JSON(http.StatusNotFound,gin.H{
			"message":"Page Not Found",
			"status":http.StatusNotFound})
	}
}

func UpdateProfile(c *gin.Context) {
	//(id,name,phone_number,contact_email,address,create_at,update_at )
	token:=c.Query("token")
	
		name:=c.PostForm("name")
		phone_number:=c.PostForm("phone_number")
		address:=c.PostForm("address")
		create_at:=c.PostForm("create_at")
		update_at:=time.Now().Format(time.RFC1123)
		avatar_id:=c.PostForm("avatar_id")
		profile:=model.Profile{}

		//(id,name,phone_number,contact_email,address,create_at,update_at
		profile.SetValueProfile("",name,phone_number,"",address,create_at,update_at)
		profile.SetAvatarID(avatar_id)
		suc,message:=controller.UpdateProfileValue(token,profile)

		if suc{
			c.JSON(http.StatusOK,gin.H{
				"message":message,
				"status":http.StatusOK})
		}else{
			c.JSON(http.StatusForbidden,gin.H{
				"message":message,
				"status":http.StatusNotFound})
		}
	
}

// func GetProfileUser(){
// 	profile_id:=c.Query("id")
// 	err,message:=controller.GetProfileWithId(profile_id)
// 	if err{
// 		c.JSON(http.StatusNotFound,gin.H{
// 			"message":message,
// 			"status":http.StatusNotFound})
// 	}
// }