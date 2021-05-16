package protocol

type FlowRequestParam struct {
	FlowID string `form:"flow_id" json:"flow_id" binding:"required"`
}
