package coc

import (
	"fmt"
	"hc-bot/api"
	"hc-bot/entity/messageUpload"
	"hc-bot/reply"
	"regexp"
	"strconv"
	"strings"
)

type Cocst struct {
}

func (cocst Cocst) ExecutePrivate(upload messageUpload.PrivateMessageUpload) {
	reply.ReplyPrivate(upload.UserId, "该命令只能在群聊中使用~")
}

func (cocst Cocst) ExecuteGroup(upload messageUpload.GroupMessageUpload) {
	reply.ReplyGroup(upload.GroupId, cocst.handle(upload.ArgStr, fmt.Sprintf("%v", upload.GroupId), fmt.Sprintf("%v", upload.UserId), upload.Sender.Nickname))
}

func (cocst Cocst) handle(argStr string, groupId string, userId string, nickname string) string {
	re := regexp.MustCompile("[a-zA-Z\u4e00-\u9fa5]{2,10}[0-9]+")
	strs := re.FindAllString(argStr, -1)
	status := make(map[string]int)
	for i := 0; i < len(strs); i++ {
		reNum := regexp.MustCompile(`\d+`)
		numStr := reNum.FindString(strs[i])
		attribute := strings.TrimRight(strs[i], numStr)
		status[attribute], _ = strconv.Atoi(numStr)
	}
	api.InsertPlayerAttributes(groupId, userId, status)
	return fmt.Sprintf("「%v」的属性设置成功", nickname)
}
