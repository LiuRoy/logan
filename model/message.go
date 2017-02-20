package model

import (
	"time"
)

type Message struct {
	MsgId uint `gorm:"column:msgid;primary_key"`
	Type string `gnorm:"column:type;type:varchar(32)"`
	InitiatorId uint `gnorm:"column:initiatorid"`
	InitiatorName string `gnorm:"column:initiatorname;type:varchar(255)"`
	InitiatorPortrait string `gnorm:"column:initiatorportrait;type:varchar(255)"`
	ConsumerId uint `gnorm:"column:consumerid"`
	ResourceId string `gnorm:"column:resource_id;type:varchar(255)"`
	ExtraInfo1 string `gnorm:"column:extra_info1;type:varchar(512)"`
	ExtraInfo2 string `gnorm:"column:extra_info2;type:varchar(512)"`
	ExtraInfo3 string `gnorm:"column:extra_info3;type:varchar(512)"`
	ExtraInfo4 string `gnorm:"column:extra_info4;type:varchar(512)"`
	InsertTime time.Time
}

func (Message) TableName() string {
	return "msgcenter_innodb"
}

func AddMessage(msgType string, initiatorId uint, initiatorName string,
	initiatorPortrait string, consumerId uint, resourceId string,
	extraInfo1, extraInfo2, extraInfo3, extraInfo4 string) *Message {
	message := Message{
		Type: msgType,
		InitiatorId: initiatorId,
		InitiatorName: initiatorName,
		InitiatorPortrait: initiatorPortrait,
		ConsumerId: consumerId,
		ResourceId: resourceId,
		ExtraInfo1: extraInfo1,
		ExtraInfo2: extraInfo2,
		ExtraInfo3: extraInfo3,
		ExtraInfo4: extraInfo4,
	}
	DbConnection.Create(&message)
	return &message
}
