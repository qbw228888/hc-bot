package coc

import (
	"fmt"
	"hc-bot/api"
	"hc-bot/entity/messageUpload"
	"hc-bot/reply"
	"strconv"
	"strings"
)

type En struct {
}

func (en En) ExecutePrivate(upload messageUpload.PrivateMessageUpload) {
	reply.ReplyPrivate(upload.UserId, en.handle(upload.ArgStr, "", strconv.FormatInt(upload.UserId, 10), upload.Sender.Nickname))
}

func (en En) ExecuteGroup(upload messageUpload.GroupMessageUpload) {
	reply.ReplyGroup(upload.GroupId, en.handle(upload.ArgStr, strconv.FormatInt(upload.GroupId, 10), strconv.FormatInt(upload.UserId, 10), upload.Sender.Nickname))
}

func (en En) handle(argStr string, groupId string, userId string, nickname string) string {
	splitArg := strings.Split(argStr, " ")
	var attr int
	if len(splitArg) == 2 {
		attributes := api.FindPlayerAttributes(groupId, userId)
		if attributes == nil {
			return "未设置角色属性值，请设置或者在命令中输入"
		}
		attrSet, ok1 := attributes[splitArg[0]]
		if ok1 == false {
			return "未设置角色属性值，请设置或者在命令中输入"
		}
		attr = attrSet
	} else if len(splitArg) == 3 {
		attrStr := splitArg[2]
		attrIn, err := strconv.Atoi(attrStr)
		if err != nil {
			return "san值请输入数字"
		}
		attr = attrIn
	} else {
		return "错误的参数，请参考help"
	}
	add := splitArg[1]
	checkResult, rand := api.GetCheckResult(attr, groupId)
	result := fmt.Sprintf("「%v」进行「%v」属性的成长检定：1d100=%v/%v", nickname, splitArg[0], rand, attr)
	switch checkResult {
	case 0, 1, 2, 3:
		result += "成长鉴定失败！无成长"
		break
	case 4, 5:
		result += fmt.Sprintf("成长检定成功！「%v」的「%v」属性增加%v", nickname, splitArg[0], api.GetLongRand(add, 1))
		break
	}
	return result
}
