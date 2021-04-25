package handler

import (
	"block_center/db"
)

func GetBlockFlow(flowId string) (blockFlow db.BlockFlow, err error){
	blockFlow, err = db.GetBlockFlowByFlowID(flowId)
	if err != nil {
		return blockFlow, err
	}

	return blockFlow, nil
}
