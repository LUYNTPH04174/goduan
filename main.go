package main

import (
	"github.com/gin-gonic/gin"
 	"goduan/route"
 	
)

func main() {
	r := gin.Default()

	r.GET("/",route.HomeDefault)
	r.GET("/login",route.GetLoginAnUser)
	rgroup:=r.Group("/app")
	{
		rgroup.POST("/post",route.InsertAnUserRouter)
	}
	
	r.Run(":8080")
}
