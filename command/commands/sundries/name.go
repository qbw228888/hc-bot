package sundries

import (
	"encoding/json"
	"fmt"
	"hc-bot/config"
	"hc-bot/entity/messageUpload"
	"hc-bot/reply"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type Name struct {
}

func (name Name) ExecutePrivate(upload messageUpload.PrivateMessageUpload) {
	reply.ReplyPrivate(upload.UserId, name.handle(upload.ArgStr))
}

func (name Name) ExecuteGroup(upload messageUpload.GroupMessageUpload) {
	reply.ReplyGroup(upload.GroupId, name.handle(upload.ArgStr))
}

type naming struct {
	Naming string `json:"naming"`
	Sex    int    `json:"sex"`
}

type nameResponse struct {
	Code   int `json:"code"`
	Result struct {
		List []naming `json:"list"`
	} `json:"result"`
}

func (name Name) handle(argStr string) string {
	var num int = 1
	var wordLen int = 2
	if argStr != "" {
		split := strings.Split(argStr, " ")
		atom, err := strconv.Atoi(split[0])
		if err != nil {
			return "错误的参数，请输入数字或不输入捏~"
		} else {
			num = atom
			if num > 50 {
				return "不能生成大于50次哦~"
			}
		}
		atom, err = strconv.Atoi(split[1])
		if err != nil {
			return "错误的参数，请输入数字或不输入捏~"
		} else {
			wordLen = atom
			if wordLen > 3 {
				return "不能生成大于3个字的名字哦~"
			}
		}
	}
	url := fmt.Sprintf("https://apis.tianapi.com/cname/index?key=%v&num=%v&wordnum=%v", config.TianJuApiKey, num, wordLen)
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	bytes, _ := io.ReadAll(resp.Body)
	nameResponse := &nameResponse{}
	err := json.Unmarshal(bytes, nameResponse)
	if err != nil {
		panic(err)
	}
	result := "随机命名如下：\n"
	arr := nameResponse.Result.List
	for i := 0; i < len(arr); i++ {
		result += fmt.Sprintf("名字：%v  ", arr[i].Naming)
		switch arr[i].Sex {
		case 1:
			result += "性别：男性\n"
			break
		case 2:
			result += "性别：女性\n"
			break
		case 3:
			result += "性别：中性\n"
		}
	}
	result = strings.TrimRight(result, "\n")
	return result
}
