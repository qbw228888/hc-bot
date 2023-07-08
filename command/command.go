package command

import (
	"hc-bot/entity/messageUpload"
	"hc-bot/reply"
)

type Command interface {
	// ExecutePrivate 执行私聊命令
	ExecutePrivate(upload messageUpload.PrivateMessageUpload)
	// ExecuteGroup 执行群聊命令
	ExecuteGroup(upload messageUpload.GroupMessageUpload)
}

type CommonCommand struct {
}

func (command CommonCommand) ExecutePrivate(upload messageUpload.PrivateMessageUpload) {
	response := "不存在的指令，请重新输入捏~"
	userId := upload.UserId
	reply.ReplyPrivate(userId, response)
}

func (command CommonCommand) ExecuteGroup(upload messageUpload.GroupMessageUpload) {
	response := "不存在的指令，请重新输入捏~"
	groupId := upload.GroupId
	reply.ReplyGroup(groupId, response)
}
