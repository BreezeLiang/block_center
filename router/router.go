package router

import (
	"block_center/config"
	"block_center/handler"
	"github.com/gin-gonic/gin"
)

func StartGin() {
	engine := gin.New()
	router := engine.Group(config.GConfig.ApiPrefix)

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.POST(
		"/get_user_flow_result",
		handler.GetBlockFlow)

	_ = engine.Run(config.GConfig.Listen.Host + ":" + config.GConfig.Listen.Port)
}
