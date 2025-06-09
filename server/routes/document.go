package routes

import (
	"termbook/handler"

	"github.com/gin-gonic/gin"
)

type DocumentRoute struct{}

func (r *DocumentRoute) Register(router *gin.Engine) {
	docRoute := router.Group("/document")
	{
		docRoute.POST("/new", handler.NewDocument)
		docRoute.POST("/open", handler.OpenDocument)
		docRoute.POST("/save", handler.SaveDocument)
		docRoute.GET("/list", handler.ListDocument)
	}
}
