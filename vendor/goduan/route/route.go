package route

import (
	"github.com/gin-gonic/gin"
	"goduan/database"
	"goduan/model"
	"net/http"
	"strconv"
	// "encoding/json"
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
	create_at:=c.PostForm("create_at")
	updated_at:=c.PostForm("updated_at")
	// create_by:=c.PostForm("create_by")
	// updated_by:=c.PostForm("create_by")



	user := model.User{}
	//SetValueUser(email,password,facebook_id,profile_id,create_at,update_at,create_by,update_by string)
	user.SetValueUser(profile_id,email, pass, facebookid,profile_id,create_at,updated_at,"","")

	// SetValueProfile(id,phone_number,contact_emal,address,create_at,update_at,create_by,update_by
	profile:=model.Profile{}
	profile.SetValueProfile(profile_id,"","","","","","","")
	user.SetProf(profile)
	err,message:=controller.InsertAnUser(user)
	if err{
		c.JSON(http.StatusOK, gin.H{
		"message":message,
		"status":  http.StatusOK})
	}else{
		c.JSON(http.StatusOK,gin.H{
		"message":message,
		"status":http.StatusOK})
	}

}

	

func GetLoginAnUser(c *gin.Context) {
	sdt := c.Query("email")
	password := c.Query("pass")
	err,user:=controller.GetAnUser(sdt, password)
	
	if err {
		c.JSON(http.StatusOK,user)
	}else{
		c.String(http.StatusNotFound,"Vui long nhap lai")
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
		"updated_at":updated_at,
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