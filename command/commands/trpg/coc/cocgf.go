package coc

import (
	"fmt"
	"hc-bot/api"
	"hc-bot/config"
	"hc-bot/entity/messageUpload"
	"hc-bot/reply"
	"strconv"
)

type Cocgf struct {
}

func (cocgf Cocgf) ExecutePrivate(upload messageUpload.PrivateMessageUpload) {
	reply.ReplyPrivate(upload.UserId, "私聊不支持此项功能")
}

func (cocgf Cocgf) ExecuteGroup(upload messageUpload.GroupMessageUpload) {
	reply.ReplyGroup(upload.GroupId, cocgf.handle(upload.ArgStr, strconv.FormatInt(upload.GroupId, 10)))
}

func (cocgf Cocgf) handle(argStr string, groupId string) string {
	if argStr == "" {
		greatFail := api.GetGreatFail(groupId)
		if greatFail == -1 {
			return fmt.Sprintf("未设置本群大失败，默认为%v", config.GreatFail)
		} else {
			return fmt.Sprintf("本群大失败为%v", greatFail)
		}
	} else {
		gf, err := strconv.Atoi(argStr)
		if err != nil {
			return "错误的参数，请输入数字捏~"
		} else {
			api.SetGreatFail(groupId, gf)
			return fmt.Sprintf("本群大失败修改成功，现在为%v", gf)
		}
	}
}
