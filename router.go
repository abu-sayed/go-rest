package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/abu-sayed/go-rest/controllers"
)

func GetRouter()  *httprouter.Router{
	router := httprouter.New()
	router.GET("/todos/:id", controllers.GetTodo)
	router.GET("/todos", controllers.GetTodos)
	router.POST("/todos", controllers.CreateTodo)
	router.PUT("/todos/:id", controllers.UpdateTodo)
	router.DELETE("/todos/:id", controllers.DeleteTodo)
	return router
}
