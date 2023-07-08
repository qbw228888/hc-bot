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

type Coc struct {
}

func (coc Coc) ExecutePrivate(upload messageUpload.PrivateMessageUpload) {
	reply.ReplyPrivate(upload.UserId, coc.handle(upload.ArgStr, upload.Sender.Nickname))
}

func (coc Coc) ExecuteGroup(upload messageUpload.GroupMessageUpload) {
	reply.ReplyGroup(upload.GroupId, coc.handle(upload.ArgStr, upload.Sender.Nickname))
}

func (coc Coc) handle(argStr string, nickname string) string {
	var times int = 1
	reNum := regexp.MustCompile("^[0-9]+$")
	if reNum.MatchString(argStr) {
		times, _ = strconv.Atoi(argStr)
	} else if argStr != "" {
		return "错误的参数，只能接收一个数字捏~"
	}
	return getCoc(nickname, times)
}

func getCoc(nickname string, times int) string {
	result := fmt.Sprintf("「%v」的调查员做成：\n", nickname)
	for i := 0; i < times; i++ {
		// 力量
		STR := api.RandN1dN2(3, 6) * 5
		// 体质
		CON := api.RandN1dN2(3, 6) * 5
		// 体型
		SIZ := api.RandN1dN2(2, 6) * 5
		// 敏捷
		DEX := api.RandN1dN2(3, 6) * 5
		// 外貌
		APP := api.RandN1dN2(3, 6) * 5
		// 智力
		INT := api.RandN1dN2(2, 6) * 5
		// 意志
		POW := api.RandN1dN2(3, 6) * 5
		// 教育
		EDU := api.RandN1dN2(2, 6) * 5
		// 幸运
		LUC := api.RandN1dN2(3, 6) * 5

		SUM := STR + CON + SIZ + DEX + APP + INT + POW + EDU + LUC
		result += fmt.Sprintf("力量：%v体质：%v体型：%v敏捷：%v外貌：%v智力：%v意志：%v教育：%v幸运：%v共计：%v/560",
			STR, CON, SIZ, DEX, APP, INT, POW, EDU, LUC, SUM)
		result += "\n"
	}
	result = strings.TrimRight(result, "\n")
	return result
}
