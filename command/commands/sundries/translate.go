package sundries

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"hc-bot/config"
	"hc-bot/entity/messageUpload"
	"hc-bot/reply"
	"io"
	"math/rand"
	"net/http"
	url2 "net/url"
	"regexp"
	"strings"
)

type Translate struct {
}

func (trans Translate) ExecutePrivate(upload messageUpload.PrivateMessageUpload) {
	reply.ReplyPrivate(upload.UserId, trans.handle(upload.ArgStr))
}

func (trans Translate) ExecuteGroup(upload messageUpload.GroupMessageUpload) {
	reply.ReplyGroup(upload.GroupId, trans.handle(upload.ArgStr))
}

func (trans Translate) handle(argStr string) string {
	reEng := regexp.MustCompile("^[a-zA-Z, ，]+$")
	if reEng.MatchString(argStr) {
		return translate(argStr, "en", "zh")
	}
	reCh := regexp.MustCompile("^[\u4e00-\u9fa5， ,]+$")
	if reCh.MatchString(argStr) {
		return translate(argStr, "zh", "en")
	}
	return "错误的参数，仅支持中英互译哦，请重新输入捏~"
}

type translateResult struct {
	Src string `json:"src"`
	Dst string `json:"dst"`
}

type translateResponse struct {
	TransResult []translateResult `json:"trans_result"`
}

func translate(query string, from string, to string) string {
	salt := rand.Int()
	sign := MD5(fmt.Sprintf("%v%v%v%v", config.BaiduTranslateAppId, query, salt, config.BaiduTranslateKey))
	url := fmt.Sprintf("https://fanyi-api.baidu.com/api/trans/vip/translate?q=%v&from=%v&to=%v&appid=%v&salt=%v&sign=%v",
		url2.QueryEscape(query), from, to, config.BaiduTranslateAppId, salt, sign)
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	bytes, _ := io.ReadAll(resp.Body)
	translateResponse := &translateResponse{}
	err := json.Unmarshal(bytes, translateResponse)
	if err != nil {
		return ""
	}
	arr := translateResponse.TransResult
	result := "以下为翻译结果：\n"
	for i := 0; i < len(arr); i++ {
		result += arr[i].Dst + "\n"
	}
	result = strings.TrimRight(result, "\n")
	return result
}

func MD5(str string) string {
	data := []byte(str) //切片
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str
}
