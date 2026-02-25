package service

import (
	"encoding/json"
	"fmt"
	"go-web/config"
	"net/http"
	"net/url"
)

type WeatherInfo struct {
	Reason      string // 返回说明
	Temperature string // 温度
	Weather     string // 天气
}

func GetWeather(city string) (WeatherInfo, error) {
	// 接口请求入参配置
	Params := url.Values{}
	Params.Set("key", config.Cfg.Api.Weather_key) // 从配置文件读取天气 API Key
	Params.Set("city", city)
	// 发起接口网络请求
	resp, err := http.Get("http://apis.juhe.cn/simpleWeather/query" + "?" + Params.Encode())
	if err != nil {
		fmt.Println("网络请求异常:", err)
		return WeatherInfo{}, err
	}
	defer resp.Body.Close()

	var responseResult map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&responseResult)
	if err != nil {
		fmt.Println("解析响应结果异常:", err)
		return WeatherInfo{}, err
	}

	// 从供应商api接口取出字段(拿自己想要的字段)
	resMap := responseResult["result"].(map[string]interface{})
	realtime := resMap["realtime"].(map[string]interface{})

	info := WeatherInfo{
		Reason:      fmt.Sprintf("%v", responseResult["reason"]),
		Temperature: fmt.Sprintf("%v", realtime["temperature"]),
		Weather:     fmt.Sprintf("%v", realtime["info"]),
	}

	return info, nil

}
