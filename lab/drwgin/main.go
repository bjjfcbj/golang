package main

import (
	"net/http"
	"time"
	"log"

	"gin"
)

func middleforadmin() gin.HandlerFunc {
	return func(c *gin.Context){
		t := time.Now()
		log.Printf("[%d] %s in %v for group admin",c.StatusCode,c.Req.RequestURI,time.Since(t))
	}
}

func main(){
	r := gin.New()
	r.Get("/",func(c *gin.Context){
		c.Html(http.StatusOK,"<h1>index</h1>")
	})

	admin := r.Group("/admin")
	admin.Use(middleforadmin())
	{
		admin.Get("/",func(c *gin.Context){
			c.Html(http.StatusOK,"<h1>hello,world</h1>")
		})
		admin.Get("/hello",func(c *gin.Context){
			c.String(http.StatusOK,"hello,you are at %s now\n",c.Path)
		})
		admin.Get("/hello/:name",func(c *gin.Context){
			c.String(http.StatusOK,"hello %s,you are at %s\n",c.Param("name"),c.Path)
		})
	}
	r.Post("/login",func(c *gin.Context){
		c.Json(http.StatusOK,gin.H{
			"username":c.PostForm("username"),
			"password":c.PostForm("password"),
		})
	})
	r.Get("/assets/*filepath",func(c *gin.Context){
		c.Json(http.StatusOK,gin.H{"filepath":c.Param("filepath")})
	})

	//the port by default:10123
	r.Run()
}
