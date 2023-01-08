package main

import (
	"QuikCMS/src"
	"QuikCMS/src/core/controllers/config"
	"QuikCMS/src/utils"
)

func main() {
	config.CheckConfig(&utils.Files{Path: "./quickCMS.json"})
	src.Run()
}
