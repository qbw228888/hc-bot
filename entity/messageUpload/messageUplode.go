package messageUpload

import (
	"encoding/json"
	"fmt"
	"hc-bot/entity"
)

type MessageUpload struct {
	entity.Upload
	MessageType string        `json:"message_type"`
	SubType     string        `json:"sub_type"`
	MessageId   int64         `json:"message_id"`
	UserId      int64         `json:"user_id"`
	Message     string        `json:"message"`
	RawMessage  string        `json:"raw_message"`
	Font        int           `json:"font"`
	Sender      entity.Sender `json:"sender"`
	StartWith   rune          // 这个命令的开始符号
	ArgStr      string        // 这个命令的参数
}

func (message *MessageUpload) JsonUnmarshal(msg []byte) {
	err := json.Unmarshal(msg, message)
	if err != nil {
		fmt.Println("error in json unmarshal: ", err)
	}
}

type PrivateMessageUpload struct {
	MessageUpload
	TargetId   int64 `json:"target_id"`
	TempSource int   `json:"temp_source"`
}

func (message *PrivateMessageUpload) JsonUnmarshal(msg []byte) {
	err := json.Unmarshal(msg, message)
	if err != nil {
		fmt.Println("error in json unmarshal: ", err)
	}
}

type GroupMessageUpload struct {
	MessageUpload
	GroupId   int64       `json:"group_id"`
	Anonymous interface{} `json:"anonymous"` //匿名消息，不是匿名消息为null
}

func (message *GroupMessageUpload) JsonUnmarshal(msg []byte) {
	err := json.Unmarshal(msg, message)
	if err != nil {
		fmt.Println("error in json unmarshal: ", err)
	}
}
