package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/abu-sayed/go-rest/controllers"
)

func GetRouter()  *httprouter.Router{
	router := httprouter.New()
	router.GET("/posts/:id", controllers.GetPost)
	router.GET("/posts", controllers.GetPosts)
	router.POST("/posts", controllers.CreatePost)
	router.PUT("/posts/:id", controllers.UpdatePost)
	router.DELETE("/posts:id", controllers.DeletePost)
	return router
}
