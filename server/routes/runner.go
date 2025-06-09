package routes

import (
	"termbook/handler"

	"github.com/gin-gonic/gin"
)

type RunnerRoute struct{}

func (r *RunnerRoute) Register(router *gin.Engine) {
	router.POST("/run-cell", handler.RunCell)
}
