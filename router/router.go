package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yuwenhui/gin-starter/controller"
	"github.com/yuwenhui/gopkg/apiutil"
)

func Initialize(r *gin.Engine) {
	r.RedirectTrailingSlash = true

	api := r.Group("/api/v1")

	// Tasks
	api.GET("/tasks", apiutil.API(controller.TaskList))
}
