package db

import (
	"block_center/common"
	"fmt"
	"github.com/pkg/errors"
)

func GetBlockFlowByFlowID(flowId string) (blockFlow BlockFlow, err error) {
	dbConn := Mysqlx
	var (
		params = make(map[string]interface{}, 0)
	)

	querySql := fmt.Sprintf(`SELECT id,user_id,flow_id,flow_type,domain,path,create_datetime as create_time FROM %s WHERE 1=1 `, common.TableBlockFlow)

	if flowId != "" {
		querySql += " AND flow_id =:flow_id "
		params["flow_id"] = flowId
	}

	{
		ns, err := dbConn.PrepareNamed(querySql)
		if err != nil {
			//log.Printf("GetBlockFlowByFlowID:", err)
			return blockFlow, errors.Wrap(err, "GetBlockFlowByFlowID:PrepareNamed")
		}
		defer ns.Close()
		err = ns.Get(&blockFlow, params)
		if err != nil {
			//log.Printf("GetBlockFlowByFlowID:", err)
			return blockFlow, errors.Wrap(err, "GetBlockFlowByFlowID:NsGet")
		}
	}

	return blockFlow, nil
}
