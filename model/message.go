package model

import (
	"time"

	"logan/config"
)

type Message struct {
	MsgId uint `gorm:"primary_key"`
	Type string `gnorm:"type:varchar(32)"`
	InitiatorId uint
	InitiatorName string `gnorm:"type:varchar(255)"`
	InitiatorPortrait string `gnorm:"type:varchar(255)"`
	ConsumerId uint
	ResourceId string `gnorm:"type:varchar(255)"`
	ExtraInfo1 string `gnorm:"type:varchar(512)"`
	ExtraInfo2 string `gnorm:"type:varchar(512)"`
	ExtraInfo3 string `gnorm:"type:varchar(512)"`
	ExtraInfo4 string `gnorm:"type:varchar(512)"`
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
	config.DbConnection.Create(&message)
	return &message
}
