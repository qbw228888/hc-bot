package sundries

import (
	"hc-bot/api"
	"hc-bot/entity/messageUpload"
	"hc-bot/reply"
	"strconv"
)

type Botoff struct {
}

func (botoff Botoff) ExecutePrivate(upload messageUpload.PrivateMessageUpload) {
	reply.ReplyPrivate(upload.UserId, "私聊不支持该命令哦~")
}

func (botoff Botoff) ExecuteGroup(upload messageUpload.GroupMessageUpload) {
	reply.ReplyGroup(upload.GroupId, botoff.handle(strconv.FormatInt(upload.GroupId, 10), upload.ArgStr, strconv.FormatInt(upload.SelfId, 10)))
}

func (botoff Botoff) handle(groupId string, argString string, botId string) string {
	if argString == botId {
		api.SetIsOn(groupId, false)
		return "Alice休眠中，等待下一次唤醒……"
	} else {
		return "请添加参数"
	}
}
