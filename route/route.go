package route

import (
	"github.com/gin-gonic/gin"
	"goduan/database"
	"goduan/model"
	"net/http"
)

var controller = database.NewUserController()

func HomeDefault(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"status":  http.StatusOK})
}

func Login(c *gin.Context) {
	firstname := c.Param("sdt")
	lastname := c.Param("pass") // shortcut for c.Request.URL.Query().Get("lastname")

	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
}

func InsertAnUserRouter(c *gin.Context) {
	sdt := c.PostForm("sdt")
	pass := c.PostForm("pass")
	facebookid := c.PostForm("facebookid")

	user := model.User{}
	user.SetValue(sdt, pass, facebookid)
	// controller:=database.NewUserController()
	controller.InsertAnUser(user)

	c.JSON(http.StatusOK, gin.H{
		"message": "Ok",
		"status":  http.StatusOK})
}

func GetLoginAnUser(c *gin.Context) {
	sdt := c.Query("sdt")
	password := c.Query("pass")
	if controller.GetAnUser(sdt, password) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"status":  http.StatusOK})
	}

}
