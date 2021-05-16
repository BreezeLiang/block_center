package handler

import (
	"block_center/protocol"
	"block_center/service"
	"block_center/util"
	"github.com/gin-gonic/gin"
	"log"
)

func GetBlockFlow(c *gin.Context) {
	reqParam := protocol.FlowRequestParam{}

	err := c.ShouldBindJSON(&reqParam)
	if err != nil {
		log.Printf(`GetBlockFlow parse param err: %v`, err)
		return
	}

	code, msg, data := service.GetBlockFlow(reqParam.FlowID)
	util.DoResponse(code, msg, data, c)
	return
}
