package coc

import (
	"fmt"
	"hc-bot/api"
	"hc-bot/entity/messageUpload"
	"hc-bot/reply"
	"strconv"
	"strings"
)

type Sc struct {
}

func (sc Sc) ExecutePrivate(upload messageUpload.PrivateMessageUpload) {
	reply.ReplyPrivate(upload.UserId, sc.handle(upload.ArgStr, "", strconv.FormatInt(upload.UserId, 10), upload.Sender.Nickname))
}

func (sc Sc) ExecuteGroup(upload messageUpload.GroupMessageUpload) {
	reply.ReplyGroup(upload.GroupId, sc.handle(upload.ArgStr, strconv.FormatInt(upload.GroupId, 10), strconv.FormatInt(upload.UserId, 10), upload.Sender.Nickname))
}

func (sc Sc) handle(argStr string, groupId string, userId string, nickname string) string {
	splitArg := strings.Split(argStr, " ")
	var san int
	if len(splitArg) == 1 {
		attributes := api.FindPlayerAttributes(groupId, userId)
		if attributes == nil {
			return "未设置角色san值，请设置或者在命令中输入"
		}
		sanSet, ok1 := attributes["san"]
		if ok1 == false {
			return "未设置角色san值，请设置或者在命令中输入"
		}
		san = sanSet
	} else {
		sanStr := splitArg[1]
		sanIn, err := strconv.Atoi(sanStr)
		if err != nil {
			return "san值请输入数字"
		}
		san = sanIn
	}
	splitMinus := strings.Split(splitArg[0], "/")
	if len(splitMinus) != 2 {
		return "错误的参数，请参考help重新输入"
	}
	checkResult, rand := api.GetCheckResult(san, groupId)
	result := fmt.Sprintf("「%v」进行sc检定：1d100=%v/%v", nickname, rand, san)
	switch checkResult {
	case 0:
		result += "大成功！san值减少：" + api.GetLongRand(splitMinus[0], 0)
		break
	case 1:
		result += "极难成功！san值减少：" + api.GetLongRand(splitMinus[0], 1)
		break
	case 2:
		result += "困难成功！san值减少：" + api.GetLongRand(splitMinus[0], 1)
		break
	case 3:
		result += "成功！san值减少：" + api.GetLongRand(splitMinus[0], 1)
		break
	case 4:
		result += "失败！san值减少：" + api.GetLongRand(splitMinus[1], 1)
		break
	case 5:
		result += "大失败！san值减少：" + api.GetLongRand(splitMinus[1], 2)
		break
	}
	return result
}
