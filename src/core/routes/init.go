package routes

import "github.com/gin-gonic/gin"

func initGroup(r *gin.Engine)  {
	group := r.Group("/init")
	{
		group.GET("")
	}
}
