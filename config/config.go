package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

var HttpRequestUrl string
var WeatherApiKey string
var BaiduTranslateAppId string
var BaiduTranslateKey string
var TianJuApiKey string
var MongoUrl string
var GreatSuccess int
var GreatFail int
var StartWith = []rune{'.', '。'}

type Config struct {
	HttpRequestUrl      string `json:"HttpRequestUrl"` // HttpRequestUrl 对应gocq监听的url
	WeatherApiKey       string `json:"WeatherApiKey"`  // WeatherApiKey 和风天气开发服务的开发者key
	BaiduTranslateAppId string `json:"BaiduTranslateAppId"`
	BaiduTranslateKey   string `json:"BaiduTranslateKey"`
	TianJuApiKey        string `json:"TianJuApiKey"`
	MongoUrl            string `json:"MongoUrl"`
	GreatSuccess        int    `json:"GreatSuccess"`
	GreatFail           int    `json:"GreatFail"`
}

func (config *Config) Init() {
	configFile, err := os.Open("./config.json")
	if err != nil {
		fmt.Println("error in open config file")
		return
	}
	defer configFile.Close()
	jsonBytes, err := io.ReadAll(configFile)
	if err != nil {
		fmt.Println("error in read config file")
		return
	}
	err = json.Unmarshal(jsonBytes, config)
	if err != nil {
		fmt.Println("error in json unmarshal config file")
		return
	}
	HttpRequestUrl = config.HttpRequestUrl
	WeatherApiKey = config.WeatherApiKey
	BaiduTranslateAppId = config.BaiduTranslateAppId
	BaiduTranslateKey = config.BaiduTranslateKey
	TianJuApiKey = config.TianJuApiKey
	MongoUrl = config.MongoUrl
	GreatSuccess = config.GreatSuccess
	GreatFail = config.GreatFail
}
