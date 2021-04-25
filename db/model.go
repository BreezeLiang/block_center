package db

type BlockFlow struct {
	Id         int    `db:"id" json:"id"`
	UserId     int64  `db:"user_id" json:"user_id"`
	FlowId     string `db:"flow_id" json:"flow_id"`
	FlowType   string `db:"flow_type" json:"flow_type"`
	Domain     string `db:"domain" json:"domain"`
	Path       string `db:"path" json:"path"`
	CreateTime string `db:"create_time" json:"create_time"`
}
