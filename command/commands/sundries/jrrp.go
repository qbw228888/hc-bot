package sundries

import (
	"fmt"
	"hc-bot/entity/messageUpload"
	"hc-bot/reply"
	"math/rand"
	"time"
)

type Jrrp struct {
}

func (jrrp Jrrp) ExecutePrivate(upload messageUpload.PrivateMessageUpload) {
	reply.ReplyPrivate(upload.UserId, jrrp.handle(upload.Sender.Nickname))
}

func (jrrp Jrrp) ExecuteGroup(upload messageUpload.GroupMessageUpload) {
	reply.ReplyGroup(upload.GroupId, jrrp.handle(upload.Sender.Nickname))
}

func (jrrp Jrrp) handle(nickname string) string {
	rand.Seed(time.Now().UnixNano())
	rp := rand.Intn(100)
	var star string
	if rp >= 1 && rp < 20 {
		star = "★☆☆☆☆"
	} else if rp >= 20 && rp < 40 {
		star = "★★☆☆☆"
	} else if rp >= 40 && rp < 60 {
		star = "★★★☆☆"
	} else if rp >= 60 && rp < 80 {
		star = "★★★★☆"
	} else {
		star = "★★★★★"
	}
	return fmt.Sprintf("今日运势：%v\n「%v」今天的好运超过%d%%的人", star, nickname, rp)
}
