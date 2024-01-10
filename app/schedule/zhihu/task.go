package zhihu

import (
	"encoding/json"
	"fmt"
	"hotinfo/app/model"
	"io/ioutil"
	"net/http"
	"time"
)

const api = "https://www.zhihu.com/api/v3/feed/topstory/hot-lists/total?limit=50&desktop=true"

func Run() {
	getInfo()
}
func getInfo() {
	// 创建HTTP客户端
	client := http.Client{}

	// 创建GET请求
	request, err := http.NewRequest("GET", api, nil)
	if err != nil {
		fmt.Println("Failed to create HTTP request:", err)
		return
	}

	// 发送请求并获取响应
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("HTTP request failed:", err)
		return
	}
	defer response.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
		return
	}

	// 解析响应内容到结构体
	var baiduHot Hot
	err = json.Unmarshal(body, &baiduHot)
	if err != nil {
		fmt.Println("Failed to unmarshal response body:", err)
		return
	}
	var data []ZhiHu
	now := time.Now().Unix()
	for _, datum := range baiduHot.Data {
		a := ZhiHu{
			UpdateVer:   now,
			Title:       datum.List.Title,
			Url:         datum.List.Url,
			Hot:         datum.DeTail_Text,
			CreatedTime: time.Now().Format("2006-01-02 15:04:05"),
			UpdatedTime: time.Now().Format("2006-01-02 15:04:05"),
		}
		data = append(data, a)
	}
	model.Conn.Table(data[0].TableName()).Create(&data)
}
