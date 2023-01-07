package handlers

import (
	"QuikCMS/src/core/controllers"
	"QuikCMS/src/utils"
	"github.com/gin-gonic/gin"
)

func CheckConfig(c *gin.Context)  {
	path := utils.Files{Path: "./quickCMS.json"}

	if !controllers.CheckConfig(&path){
		c.String(404, "Incorrect")
	}

	c.String(200, "Correct")
}