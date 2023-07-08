package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
	"hc-bot/config"
	"hc-bot/dispatcher"
	"hc-bot/entity/messageUpload"
)

func main() {
	r := gin.Default()
	m := melody.New()

	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	var config = &config.Config{}
	config.Init()

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		var messageMap map[string]interface{}
		errToMap := json.Unmarshal(msg, &messageMap)
		if errToMap != nil {
			fmt.Println("JSON转换失败：", errToMap)
			return
		}
		// 不处理非消息的信息
		if messageMap["post_type"] != "message" {
			return
		}
		// 非命令消息不处理
		message := messageUpload.MessageUpload{}
		message.JsonUnmarshal(msg)
		flag, rune := isOrder(message)
		if flag == false {
			return
		}
		// 分配处理群聊和私聊消息，直接开一个协程进行处理，不影响接收下一个消息
		if messageMap["message_type"] == "group" {
			message := messageUpload.GroupMessageUpload{}
			message.JsonUnmarshal(msg)
			message.StartWith = rune
			go dispatcher.DispatchGroup(message)
		} else {
			message := messageUpload.PrivateMessageUpload{}
			message.JsonUnmarshal(msg)
			message.StartWith = rune
			go dispatcher.DispatchPrivate(message)
		}

	})

	r.Run(":2288")
}

// 判断是否是命令, 并且返回这个命令的起始字符
func isOrder(message messageUpload.MessageUpload) (bool, rune) {
	char := []rune(message.Message)[0]
	for i := 0; i < len(config.StartWith); i++ {
		if char == config.StartWith[i] {
			return true, char
		}
	}
	return false, 0
}
