package main

import (
	"httpServer/conf"
	"httpServer/http/get"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()
	initRouter(g)
	err := g.Run(conf.GetHost())
	if err != nil {
		panic(err)
	}
}

// 初始化路由
func initRouter(engine *gin.Engine) {
	initConfig(engine)
	initGetRouter(engine)
	initPostRouter(engine)
}

func initConfig(engine *gin.Engine) {
	engine.LoadHTMLGlob("./website/web/html/*")
	engine.Static("/static", "./website/web/static")
}

// 初始化Get路由
func initGetRouter(engine *gin.Engine) {
	engine.GET("/", func(ctx *gin.Context) {
		get.GetIndexPage(ctx)
	})
}

// 初始化Post路由
func initPostRouter(engine *gin.Engine) {

}
