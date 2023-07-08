package sundries

import (
	"encoding/json"
	"fmt"
	"hc-bot/config"
	"hc-bot/entity/messageUpload"
	"hc-bot/reply"
	"io"
	"net/http"
	url2 "net/url"
)

type Weather struct {
}

func (weather Weather) ExecutePrivate(upload messageUpload.PrivateMessageUpload) {
	reply.ReplyPrivate(upload.UserId, weather.handle(upload.ArgStr))

}

func (weather Weather) ExecuteGroup(upload messageUpload.GroupMessageUpload) {
	reply.ReplyGroup(upload.GroupId, weather.handle(upload.ArgStr))
}

func (weather Weather) handle(argStr string) string {
	if argStr == "" {
		return "错误的参数，请重新输入捏~"
	}
	cityId := getLocationId(argStr)
	if cityId == "" {
		return "没找到这个城市，换一个吧~"
	}
	return getWeatherNow(argStr, cityId)
}

// 封装城市信息
type cityInfo struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

// 封装搜寻城市id的回复
type locationResponse struct {
	Code     string     `json:"code"`
	Location []cityInfo `json:"location"`
}

// 根据城市名请求城市id
func getLocationId(city string) string {
	// 注意需要对可能包含中文的参数进行转码
	url := fmt.Sprintf("https://geoapi.qweather.com/v2/city/lookup?location=%v&key=%v", url2.QueryEscape(city), config.WeatherApiKey)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	response := &locationResponse{}
	err = json.Unmarshal(b, response)
	if err != nil {
		panic(err)
	}
	if response.Code != "200" {
		return ""
	}
	return response.Location[0].Id
}

type weatherNowResponse struct {
	Code       string `json:"code"`
	UpdateTime string `json:"updateTime"`
	Now        struct {
		Temp      string `json:"temp"`
		FeelsLike string `json:"feelsLike"`
		Text      string `json:"text"`
		WindDir   string `json:"windDir"`
		WindScale string `json:"windScale"`
		Humidity  string `json:"humidity"`
		Vis       string `json:"vis"`
	} `json:"now"`
}

func getWeatherNow(city string, cityId string) string {
	url := fmt.Sprintf("https://devapi.qweather.com/v7/weather/now?location=%v&key=%v", cityId, config.WeatherApiKey)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	weatherNowResponse := &weatherNowResponse{}
	err = json.Unmarshal(b, weatherNowResponse)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("今天%v天气为%v\n气温%v度\n体感温度%v度\n刮%v，风力%v级\n相对湿度%v\n可见度%v公里\n更新时间%v",
		city, weatherNowResponse.Now.Text, weatherNowResponse.Now.Temp, weatherNowResponse.Now.FeelsLike,
		weatherNowResponse.Now.WindDir, weatherNowResponse.Now.WindScale, weatherNowResponse.Now.Humidity,
		weatherNowResponse.Now.Vis, weatherNowResponse.UpdateTime)
}
