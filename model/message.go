package model

import (
	"time"
)

type Message struct {
	MsgId uint `gorm:"column:msgid;primary_key"`
	Type string `gorm:"column:type;type:varchar(32)"`
	InitiatorId uint `gorm:"column:initiatorid"`
	InitiatorName string `gorm:"column:initiatorname;type:varchar(255)"`
	InitiatorPortrait string `gorm:"column:initiatorportrait;type:varchar(255)"`
	ConsumerId uint `gorm:"column:consumerid"`
	ResourceId string `gorm:"column:resource_id;type:varchar(255)"`
	ExtraInfo1 string `gorm:"column:extra_info1;type:varchar(512)"`
	ExtraInfo2 string `gorm:"column:extra_info2;type:varchar(512)"`
	ExtraInfo3 string `gorm:"column:extra_info3;type:varchar(512)"`
	ExtraInfo4 string `gorm:"column:extra_info4;type:varchar(512)"`
	InsertTime time.Time `gorm:"column:insert_time" sql:"DEFAULT:current_timestamp"`
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
	DbConnection.NewRecord(message)
	return &message
}

func GetMessage(messageId uint) *Message {
	message := Message{}
	DbConnection.Where("msgid = ?", messageId).First(&message)
	return &message
}
