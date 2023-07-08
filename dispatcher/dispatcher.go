package dispatcher

import (
	"hc-bot/api"
	"hc-bot/command"
	"hc-bot/entity/messageUpload"
	"hc-bot/reply"
	"strconv"
	"strings"
)

func DispatchPrivate(messageUpload messageUpload.PrivateMessageUpload) {
	message := messageUpload.Message
	for k, v := range command.Commands {
		if f, arg := isCommand(message, k, messageUpload.StartWith); f {
			messageUpload.ArgStr = arg
			v.ExecutePrivate(messageUpload)
			return
		}
	}
	command.CommonCommand{}.ExecutePrivate(messageUpload)
}

func DispatchGroup(messageUpload messageUpload.GroupMessageUpload) {
	message := messageUpload.Message
	if b, _ := isCommand(message, "boton", messageUpload.StartWith); b {
		api.SetIsOn(strconv.FormatInt(messageUpload.GroupId, 10), true)
		reply.ReplyGroup(messageUpload.GroupId, "Alice启动，为您服务……")
		return
	}
	if api.GetIsOn(strconv.FormatInt(messageUpload.GroupId, 10)) == false {
		return
	}
	for k, v := range command.Commands {
		if f, arg := isCommand(message, k, messageUpload.StartWith); f {
			messageUpload.ArgStr = arg
			v.ExecuteGroup(messageUpload)
			return
		}
	}
	command.CommonCommand{}.ExecuteGroup(messageUpload)
}

// 判断message是否是这个命令，如果是，添加这个命令的参数
func isCommand(message string, command string, startWith rune) (bool, string) {
	f := strings.HasPrefix(message, string(startWith)+command+" ") || (strings.LastIndex(message, string(startWith)+command) != -1 && strings.LastIndex(message, string(startWith)+command) == len(message)-len(string(startWith)+command))
	if f {
		arr := strings.SplitN(message, " ", 2)
		if len(arr) < 2 {
			return true, ""
		} else {
			return true, arr[1]
		}
	} else {
		return false, ""
	}
}
