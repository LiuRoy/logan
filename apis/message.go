package apis

import (
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"

	"logan/model"
)

type message struct {
	Type string `json:"type" binding:"required"`
	InitiatorId uint `json:"initiator" binding:"required"`
	ConsumerId uint `json:"consumer" binding:"required"`
	ResourceId string `json:"resource_id"`
	IsFollow bool `json:"isfollow"`
	Gcid string `json:"gcid"`
	Cid uint `json:"cid"`
	Response string `json:"response"`
	reply string `json:"reply"`
}

func bool2string(b bool) string {
	if b {
		return "1"
	} else {
		return "0"
	}
}

func AddMessage(c *gin.Context) {
	var param message
	if c.BindJSON(&param) == nil {
		switch param.Type {
		case "follow":
			model.AddMessage(
				param.Type, param.InitiatorId,  "aaa",
				"bbb",  param.ConsumerId, "",
				bool2string(param.IsFollow), "", "", "")
		default:
			panic("unknown type")
		}
		c.JSON(http.StatusOK, gin.H{"result": "ok"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"result": "bad params"})
	}
}

func GetMessage(c *gin.Context) {
	messageId, exist := c.GetQuery("message_id")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"result": "bad params"})
	} else {
		msgId, _ := strconv.ParseUint(messageId, 10, 0)
		message := model.GetMessage(uint(msgId))
		c.JSON(http.StatusOK, gin.H{"result": "ok", "message": *message})
	}
}
