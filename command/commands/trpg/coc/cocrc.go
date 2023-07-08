package coc

import (
	"fmt"
	"hc-bot/api"
	"hc-bot/entity/messageUpload"
	"hc-bot/reply"
	"regexp"
	"strconv"
)

type Cocrc struct {
}

func (cocrc Cocrc) ExecutePrivate(upload messageUpload.PrivateMessageUpload) {
	reply.ReplyPrivate(upload.UserId, "私聊不支持此项功能")
}

func (cocrc Cocrc) ExecuteGroup(upload messageUpload.GroupMessageUpload) {
	reply.ReplyGroup(upload.GroupId, cocrc.handle(upload.ArgStr, strconv.FormatInt(upload.GroupId, 10), strconv.FormatInt(upload.UserId, 10), upload.Sender.Nickname))
}

func (cocrc Cocrc) handle(argStr string, groupId string, userId string, nickname string) string {
	re := regexp.MustCompile("[a-zA-Z\u4e00-\u9fa5]{2,10}")
	if !re.MatchString(argStr) {
		return "错误的参数，请重新输入~"
	}
	arg := re.FindString(argStr)
	attributes := api.FindPlayerAttributes(groupId, userId)
	if attributes == nil {
		return "没有设定调查员参数~"
	}
	attribute, ok := attributes[arg]
	if ok == false {
		return "没有设置这项属性"
	} else {
		checkResult, rand := api.GetCheckResult(attribute, groupId)
		result := fmt.Sprintf("「%v」进行「%v」检定：1d100=%v/%v，", nickname, arg, rand, attribute)
		switch checkResult {
		case 0:
			result += "大成功！「祂盯上你了」"
			break
		case 1:
			result += "极难成功！「幸运女神在微笑」"
			break
		case 2:
			result += "困难成功！「超常发挥」"
			break
		case 3:
			result += "成功！「侥幸而已」"
			break
		case 4:
			result += "失败！「不要相信概率，专家也会失手」"
			break
		case 5:
			result += "大失败！「深渊在看着你！」"
			break
		}
		return result
	}
}
