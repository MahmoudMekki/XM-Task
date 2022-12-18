package models

const (
	ActivityLogTableName = "activity_logs"
	ActivityLogTopic     = "activity_logs"
)

type Log struct {
	Id       int64  `gorm:"column:id;primary_key;auto_increment" json:"id"`
	UserId   int64  `gorm:"column:user_id;not null" json:"user_id"`
	Activity string `gorm:"column:activity;not null" json:"activity"`
	Data     []byte `gorm:"column:data;not null" json:"data"`
}

func (Log) TableName() string {
	return ActivityLogTableName
}
