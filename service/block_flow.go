package service

import (
	"block_center/db"
	"log"
)

func GetBlockFlow(flowId string) (code int, msg string, data interface{}) {
	blockFlow, err := db.GetBlockFlowByFlowID(flowId)
	if err != nil {
		log.Printf(`GetBlockFlow service err:%v`, err)
		return 0, "获取失败", blockFlow
	}

	return 0, "获取成功", blockFlow
}
