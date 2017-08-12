package main

import (
	"github.com/gin-gonic/gin"
	"goduan/route"
	"net/http"
	"time"
	// "os"
)

func main() {
	// port:=os.Getenv("PORT")

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/", route.HomeDefault)
	r.GET("/login", route.GetLoginAnUser)
	r.GET("/categorys", route.GetCategorys)
	r.GET("profile",route.GetProfileUser)
	r.GET("/list",route.GetListDetailByCategory)
	r.GET("/find",route.SearchPost)
	rgroup := r.Group("/api")
	{
		rgroup.POST("/newuser", route.InsertAnUserRouter)
		rgroup.POST("/newcate", route.InsertACategory)
		rgroup.POST("/upprofile",route.UpdateProfile)
		rgroup.POST("/post",route.InsertAPostDetail)
		rgroup.POST("/feed",route.InsertAFeedBack)
	}

	 s := &http.Server{
        Addr:           ":8080",
        Handler:        r,
        ReadTimeout:    5 * time.Second,
        WriteTimeout:   5 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }
    s.ListenAndServe()
}