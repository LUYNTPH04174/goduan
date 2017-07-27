package main

import (
	"github.com/gin-gonic/gin"
	"goduan/route"
	// "os"
)

func main() {
	// port:=os.Getenv("PORT")

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/login", route.GetLoginAnUser)
	r.GET("/categorys", route.GetCategorys)
	r.GET("profile",route.GetProfileUser)
	r.GET("/",route.GetListDetailByCategory)
	rgroup := r.Group("/api")
	{
		rgroup.POST("/newuser", route.InsertAnUserRouter)
		rgroup.POST("/newcate", route.InsertACategory)
		rgroup.POST("/upprofile",route.UpdateProfile)
		rgroup.POST("/post",route.InsertAPostDetail)
	}

	// r.Run(":"+port)
	r.Run(":8080")
}