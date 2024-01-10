package hyper

import (
	"encoding/json"
	"hotinfo/app/model"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const api = "https://cache.thepaper.cn/contentapi/wwwIndex/rightSidebar"

func Run() {
	getHyper()
}
func getHyper() {
	// 发送 HTTP 请求
	resp, err := http.Get(api)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// 解析 JSON 数据
	var responseData map[string]interface{}
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		log.Fatal(err)
	}

	// 提取 hotNews 数组
	data, ok := responseData["data"].(map[string]interface{})
	if !ok {
		log.Fatal("提取数据失败")
	}
	hotNewsArray, ok := data["hotNews"].([]interface{})
	if !ok {
		log.Fatal("Failed to extract 'hotNews' array from JSON")
	}
	// 遍历 hotNews 数组并提取字段
	for _, newsItem := range hotNewsArray {
		if newsMap, ok := newsItem.(map[string]interface{}); ok {
			// 提取 contId 字段
			var name, tag string
			if contId, ok := newsMap["contId"].(string); ok {
				// 拼接 URL
				url := "https://www.thepaper.cn/newsDetail_forward_" + contId
				// 提取 name 字段
				if nameVal, ok := newsMap["name"].(string); ok {
					name = nameVal
				}
				// 提取 tagList 数组
				tagList, ok := newsMap["tagList"].([]interface{})
				if !ok {
					log.Fatal("提取 tag 失败")
				}
				// 遍历 tagList 数组并提取第一个 tag
				for _, tagItem := range tagList {
					if tagMap, ok := tagItem.(map[string]interface{}); ok {
						if tagVal, ok := tagMap["tag"].(string); ok {
							tag = tagVal
							break // 只提取第一个 tag，可以直接退出内层循环
						}
					}
				}
				tmp := Hyper{
					UpdateVer:   time.Now().Unix(),
					Title:       name,
					KeyWord:     tag,
					Url:         url,
					CreatedTime: time.Now(),
					UpdatedTime: time.Now(),
				}
				model.Conn.Create(&tmp)
			}
		}
	}
}
