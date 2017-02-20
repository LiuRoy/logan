package apis

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"logan/model"
)

type message struct {
	Type string `json:"type" bind:"required"`
	InitiatorId uint `json:"initiator" bind:"required"`
	ConsumerId uint `json:"consumer" bind:"required"`
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
