package reply

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hc-bot/config"
	"net/http"
)

func ReplyPrivate(userId int64, message string) {
	data := map[string]interface{}{
		"user_id":     userId,
		"message":     message,
		"auto_escape": false,
		"group_id":    nil,
	}
	json, err := json.Marshal(data)
	if err != nil {
		fmt.Println("error in marshal to json:", err)
		return
	}
	http.Post(config.HttpRequestUrl+"send_private_msg", "application/json", bytes.NewReader(json))
}

func ReplyGroup(groupId int64, message string) {
	data := map[string]interface{}{
		"group_id":    groupId,
		"message":     message,
		"auto_escape": false,
	}
	json, err := json.Marshal(data)
	if err != nil {
		fmt.Println("error in marshal to json:", err)
		return
	}
	http.Post(config.HttpRequestUrl+"send_group_msg", "application/json", bytes.NewReader(json))
}
