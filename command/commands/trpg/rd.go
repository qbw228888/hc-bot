package trpg

import (
	"fmt"
	"hc-bot/api"
	"hc-bot/entity/messageUpload"
	"hc-bot/reply"
)

type Rd struct {
}

func (rd Rd) ExecutePrivate(upload messageUpload.PrivateMessageUpload) {
	reply.ReplyPrivate(upload.UserId, fmt.Sprintf("「%v」掷骰：", upload.Sender.Nickname)+rd.handle(upload.ArgStr))
}

func (rd Rd) ExecuteGroup(upload messageUpload.GroupMessageUpload) {
	reply.ReplyGroup(upload.GroupId, fmt.Sprintf("「%v」掷骰：", upload.Sender.Nickname)+rd.handle(upload.ArgStr))
}

func (rd Rd) handle(argStr string) string {
	return api.GetLongRand(argStr, 1)
}

type Rh struct {
	*Rd
}

func (rh Rh) ExecutePrivate(upload messageUpload.PrivateMessageUpload) {
	reply.ReplyPrivate(upload.UserId, "私聊不支持该功能~")
}

func (rh Rh) ExecuteGroup(upload messageUpload.GroupMessageUpload) {
	reply.ReplyPrivate(upload.UserId, fmt.Sprintf("「%v」在群「%v」掷暗骰：", upload.Sender.Nickname, upload.GroupId)+rh.handle(upload.ArgStr))
	reply.ReplyGroup(upload.GroupId, fmt.Sprintf("「%v」偷偷地掷骰", upload.Sender.Nickname))
}
