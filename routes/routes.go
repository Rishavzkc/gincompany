package routes

import (
	"final/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {

	comp := router.Group("company")
	{
		comp.GET("/", controllers.ShowAll)
		comp.GET("/", controllers.ShowAll)
		comp.GET("/:id", controllers.Show)
		comp.POST("/", controllers.Create)
		comp.PUT("/", controllers.Update)
		comp.DELETE("/:id", controllers.Delete)
	}

	return router
}
