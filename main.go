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

	r.GET("/", route.HomeDefault)
	r.GET("/login", route.GetLoginAnUser)
	r.GET("/categorys", route.GetCategorys)
	rgroup := r.Group("/api")
	{
		rgroup.POST("/newuser", route.InsertAnUserRouter)
		rgroup.POST("/newcate", route.InsertACategory)
	}

	// r.Run(":"+port)
	r.Run(":8080")
}