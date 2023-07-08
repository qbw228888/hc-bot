package sundries

import (
	"hc-bot/entity/messageUpload"
	"hc-bot/reply"
)

type Help struct {
}

func (help Help) ExecutePrivate(upload messageUpload.PrivateMessageUpload) {
	userId := upload.UserId
	response := help.handle(upload.ArgStr)
	reply.ReplyPrivate(userId, response)
}

func (help Help) ExecuteGroup(upload messageUpload.GroupMessageUpload) {
	groupId := upload.GroupId
	response := help.handle(upload.ArgStr)
	reply.ReplyGroup(groupId, response)
}

func (help Help) handle(argStr string) string {
	var response string
	switch argStr {
	case "":
		response = responseNoArg()
		break
	case "跑团":
		response = responseForTrpg()
		break
	case "杂类":
		response = responseForSundry()
		break
	default:
		response = responseErrorArg()
		break
	}
	return response
}

func responseNoArg() string {
	result := "骰娘Alice v1.0\n使用帮助：\n------------\n命令以「.」或「。」开头\n注意命令与参数的空格\n------------\n"
	result += "1.跑团类\n（包括简易车卡、掷骰子）\n参考命令：.help 跑团\n"
	result += "2.杂类\n（包括天气、今日人品等）\n参考命令：.help 杂类"
	return result
}

func responseForTrpg() string {
	response := "跑团使用帮助：\n------------\n"
	response += "1.coc人物做成\n命令：.coc 次数\n「这张卡又会在多久之后被撕掉呢？」\n------------\n"
	response += "2.掷骰子\n示例命令：.rd 4d6+2d3\n参数不写默认为1d100\n也支持参数为纯数字，如.rd 10\n暗骰命令为.rh，格式如前所述，暗骰前请先加好友\n「也许你不相信，但骰娘已经计划好了一切」\n------------\n"
	response += "3.记录coc人物\n示例命令：.st 侦查10聆听10\n参数不得为空\n支持空白人物表的骰娘输入格式\n「再把技能点的高点？这并不影响你步入疯狂……」\n------------\n"
	response += "4.进行coc检定\n示例命令：.rc 侦查\n参数为一个技能，不得为空\n「✞向骰娘祈祷吧！骰门✞」\n------------\n"
	response += "5.进行大成功的设置或者查看当前群的大成功数值\n示例命令：.cocgs\n无参数查询当前群的大成功数值，有参数设置当前群的大成功数值\n「大成功？最好是吧」\n------------\n"
	response += "6.进行大失败的设置或者查看当前群的大失败数值\n示例命令：.cocgf\n规则与大成功一样\n「大失败？最好是吧」\n------------\n"
	response += "7.进行sc检定\n示例命令：.sc 1d3/2d3+1 60\n上述命令代表san值为60，成功检定san值减少1d3，失败减少2d3+1\n如果使用st命令在群内设置过san属性，则可以省略san值参数，如果是私聊不可省略\n「你的意志正在接受考验」\n------------\n"
	response += "8.进行成长检定\n示例命令：.en 教育 1d10 60\n上述命令代表为60的教育进行一次1d10的成长\n如果用st命令设置过属性则可以省略60这个参数，如果在私聊中使用则不可以省略\n「进步，成长……但人类总是有极限……」\n------------\n"
	response += "关于dnd的更多功能仍在开发中……"
	return response
}

func responseForSundry() string {
	response := "杂类使用帮助：\n------------\n"
	response += "1.世界天气\n示例命令：.weather 伊拉克\n「世界上总有一个地方是正午」\n------------\n"
	response += "2.中英互译，不要输入逗号以外的标点哦\n示例命令：.trans 雨季\n「你真的愿意相信百度翻译？」\n------------\n"
	response += "3.今日人品\n示例命令：.jrrp\n「相信的人不会多测」\n------------\n"
	response += "4.随机名字，一次最多生成50个，只有中文，第一个参数为次数，第二个为名字字数(字数最多为3)\n示例命令：.name 3\n「生成出来的玩意可能狗屁不通」\n------------"
	return response
}

func responseErrorArg() string {
	return "输入了错误的参数捏~，请参考如下help命令提示：\n" + responseNoArg()
}
