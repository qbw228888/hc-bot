package coc

import (
	"fmt"
	"hc-bot/api"
	"hc-bot/config"
	"hc-bot/entity/messageUpload"
	"hc-bot/reply"
	"strconv"
)

type Cocgs struct {
}

func (cocgs Cocgs) ExecutePrivate(upload messageUpload.PrivateMessageUpload) {
	reply.ReplyPrivate(upload.UserId, "私聊不支持此项功能")
}

func (cocgs Cocgs) ExecuteGroup(upload messageUpload.GroupMessageUpload) {
	reply.ReplyGroup(upload.GroupId, cocgs.handle(upload.ArgStr, strconv.FormatInt(upload.GroupId, 10)))
}

func (cocgs Cocgs) handle(argStr string, groupId string) string {
	if argStr == "" {
		greatSuccess := api.GetGreatSuccess(groupId)
		if greatSuccess == -1 {
			return fmt.Sprintf("未设置本群大成功，默认为%v", config.GreatSuccess)
		} else {
			return fmt.Sprintf("本群大成功为%v", greatSuccess)
		}
	} else {
		gs, err := strconv.Atoi(argStr)
		if err != nil {
			return "错误的参数，请输入数字捏~"
		} else {
			api.SetGreatSuccess(groupId, gs)
			return fmt.Sprintf("本群大成功修改成功，现在为%v", gs)
		}
	}
}
