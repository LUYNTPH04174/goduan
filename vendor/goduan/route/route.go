package route

import (
	"github.com/gin-gonic/gin"
	"goduan/database"
	"goduan/model"
	"net/http"
	"strconv"
	// "encoding/json"
	"time"
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
	profile_id:=strconv.Itoa(controller.GetUsersCount()+1)
	create_at:=time.Now().Format(time.RFC1123)
	name:=c.PostForm("name")
	user := model.User{}
	//SetValueUser(email,password,facebook_id,profile_id,create_at string)
	user.SetValueUser(profile_id,email, pass, facebookid,profile_id,create_at)

	// SetValueProfile(id,name,phone_number,contact_emal,address,create_at,update_at)
	profile:=model.Profile{}
	profile.SetValueProfile(profile_id,name,"","","",create_at,"")
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
	succ,user:=controller.GetAnUser(email, password)
	succ2,pro:=controller.GetProfileWithUser(user.Profile_id)
	if succ &&succ2{
		c.JSON(http.StatusOK,gin.H{
			"user":user,
			"profile":pro})
	}else{
		c.String(http.StatusNotFound,"Thông tin chưa chính xác")
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
		c.String(http.StatusNotFound,"Page Not Found")
	}
}

func UpdateProfile(c *gin.Context) {
	
}