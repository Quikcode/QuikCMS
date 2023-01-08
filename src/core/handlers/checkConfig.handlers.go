package handlers

import (
	"QuikCMS/src/core/controllers/config"
	"QuikCMS/src/utils"
	"github.com/gin-gonic/gin"
)

func CheckConfig(c *gin.Context) {
	path := utils.Files{Path: "./quickCMS.json"}

	if !config.CheckConfig(&path) {
		c.String(404, "Incorrect")
	}

	c.String(200, "Correct")
}
