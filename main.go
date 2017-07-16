package main

import (
	"github.com/gin-gonic/gin"
	"goduan/route"
	// "os"
)

func main() {
	// port:=os.Getenv("PORT")
	r := gin.Default()

	r.GET("/", route.HomeDefault)
	r.GET("/login", route.GetLoginAnUser)
	rgroup := r.Group("/app")
	{
		rgroup.POST("/post", route.InsertAnUserRouter)
	}

	// r.Run(":"+port)
	r.Run(":8080")
}
