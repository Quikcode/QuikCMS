package routes

import (
	"QuikCMS/src/core/handlers"
	"QuikCMS/src/vars"
	"github.com/gin-gonic/gin"
)

func Set(r *gin.Engine)  {
	r.POST("check_config", handlers.CheckConfig)

	initGroup(r)

	r.NoRoute(func(c *gin.Context) {
		c.File(vars.PathDistUI + "/index.html")
	})
}