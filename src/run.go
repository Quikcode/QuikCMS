package src

import (
	"QuikCMS/src/core/routes"
	"QuikCMS/src/vars"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()

	r.Use(static.Serve("/", static.LocalFile(vars.PathDistUI, true)))
	routes.Set(r)

	err := r.Run("0.0.0.0:80")
	if err != nil {
		panic(err)
	}
}