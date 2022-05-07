package routes

import (
	"gin-api-rest/controller"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/students", controller.GetAll)
	r.GET("/:name", controller.Greeting)
	r.GET("/students/:id", controller.GetOneByID)
	r.GET("/students/cpf/:cpf", controller.GetOneByCpf)
	r.POST("/students", controller.Create)
	r.DELETE("/students/:id", controller.Delete)
	r.PATCH("/students/:id", controller.Update)
	r.GET("/index", controller.ShowIndexPage)
	r.NoRoute(controller.PageNotFound)
	r.Run()
}
