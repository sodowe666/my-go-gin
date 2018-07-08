package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"web/config"
	_ "web/config/core"
	"web/router"
)

func main() {
	//判断环境，开启日志
	if config.GetConfig().Server.Environment == "dev" {
		os.MkdirAll("./gin-log", 0777)
		// Disable Console Color, you don't need console color when writing the logs to file.
		gin.DisableConsoleColor()
		// Logging to a file.
		f, _ := os.Create("./gin-log/info.log")
		gin.DefaultWriter = io.MultiWriter(f)
	}
	engine := router.InitRouter()
	engine.Run()
}
