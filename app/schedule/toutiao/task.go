package toutiao

import (
	"encoding/json"
	"fmt"
	"hotinfo/app/model"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const api = "https://www.toutiao.com/hot-event/hot-board/?origin=toutiao_pc&_signature=_02B4Z6wo00d01fz2EaAAAIDCqXSaPLKo-Pn80hUAABqj2McKKyN-ZFKSeYnDCzFmiHfyLZVJUtrSCHyGtp2v7Ztb7pLyecXTYgXFN6.XJUqt0-qtlRw2HTe2N25GxfGIYvmXAjOtTfj2fexBfc"

func Run() {
	ticker := time.NewTicker(5 * time.Minute)
	defer func() {
		ticker.Stop()
	}()

	for {
		select {
		case <-ticker.C:
			getInfo()
		}
	}
}
func Do() {
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
	var data []TouTiao
	now := time.Now().Unix()
	for _, datum := range baiduHot.List {
		a := TouTiao{
			Icon:        datum.Label,
			LabelDesc:   datum.LabelDesc,
			UpdateVer:   now,
			Title:       datum.Title,
			Url:         datum.Url,
			Hot:         datum.Hot,
			CreatedTime: time.Now().Format("2006-01-02 15:04:05"),
			UpdatedTime: time.Now().Format("2006-01-02 15:04:05"),
		}
		data = append(data, a)
	}
	//fmt.Println(data)
	model.Conn.Table(data[0].TableName()).Create(&data)
}
func Refresh() []TouTiao {
	var maxUpdateVer int64

	// 查询最大的 update_ver
	result := model.Conn.Model(&TouTiao{}).Select("MAX(update_ver) as max_update_ver").Scan(&maxUpdateVer)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	// 查询所有 update_ver 为最大值的记录
	var toutiaoList []TouTiao
	result = model.Conn.Where("update_ver = ?", maxUpdateVer).Find(&toutiaoList)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	// 打印查询结果
	fmt.Printf("Data with max update_ver (%d):\n", maxUpdateVer)
	return toutiaoList

}
