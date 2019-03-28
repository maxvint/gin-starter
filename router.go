package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yuwenhui/gin-starter/controller"
	"github.com/yuwenhui/gopkg/apiutil"
)

func CreateRouter(r *gin.Engine) {
	r.RedirectTrailingSlash = true

	api := r.Group("/api/v1")
	api.GET("/tasks", apiutil.API(controller.TaskList))
}
