package zhihu

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"hotinfo/app/model"
	"hotinfo/app/tools"
	"io/ioutil"
	"net/http"
	"time"
)

const api = "https://www.zhihu.com/api/v3/feed/topstory/hot-lists/total?limit=50&desktop=true"

func Run() {
	ticker := time.NewTicker(1 * time.Minute)
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
		logrus.Error("zhihu:Failed to create HTTP request:", err)
		//fmt.Println("Failed to create HTTP request:", err)
		return
	}

	// 发送请求并获取响应
	response, err := client.Do(request)
	if err != nil {
		logrus.Error("zhihu:HTTP request failed:", err)
		//fmt.Println("HTTP request failed:", err)
		return
	}
	defer response.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logrus.Error("zhihu:Failed to read response body:", err)
		//fmt.Println("Failed to read response body:", err)
		return
	}

	// 解析响应内容到结构体
	var baiduHot Hot
	err = json.Unmarshal(body, &baiduHot)
	if err != nil {
		logrus.Error("zhihu:Failed to unmarshal response body:", err)
		//fmt.Println("Failed to unmarshal response body:", err)
		return
	}
	var data []ZhiHu
	now := time.Now().Unix()
	var hotinfoStr string
	for _, datum := range baiduHot.Data {
		a := ZhiHu{
			UpdateVer:   now,
			Title:       datum.List.Title,
			Url:         datum.List.Url,
			Hot:         datum.DeTail_Text,
			CreatedTime: time.Now(),
			UpdatedTime: time.Now(),
		}
		data = append(data, a)
		hotinfoStr = hotinfoStr + datum.List.Title + datum.List.Url
	}
	hashStr := tools.Sha256Hash(hotinfoStr)

	value, err := model.RedisClient.Get(context.Background(), "zhihu").Result()
	if err == redis.Nil {
		err = model.RedisClient.Set(context.Background(), "zhihu_hot", hashStr, 0).Err()
		if err != nil {
			logrus.Error("zhihu:Failed to set value in Redis:", err)
			//fmt.Printf("Failed to set value in Redis: %v", err)
			return

		}
		model.Conn.Create(data)
	} else if err != nil {
		logrus.Error("zhihu:Error getting value from Redis:", err)
		//fmt.Printf("Error getting value from Redis: %v", err)
	} else {

		if hashStr != value {
			err = model.RedisClient.Set(context.Background(), "zhihu_hot", hashStr, 0).Err()
			if err != nil {
				logrus.Error("zhihu:Error setting value from Redis:", err)
			}
			err = model.Conn.Create(data).Error
			if err != nil {
				logrus.Error("zhihu:db_create:", err)
			}
		} else {
			var maxUpdateVer int64
			var updateSlice []ZhiHu
			model.Conn.Model(&ZhiHu{}).Select("MAX(update_ver) as max_update_ver").Scan(&maxUpdateVer)
			model.Conn.Where("update_ver = ?", maxUpdateVer).Find(&updateSlice)
			for _, record := range updateSlice {
				record.UpdateVer = now
				record.UpdatedTime = time.Now()
				err := model.Conn.Save(&record).Error
				if err != nil {
					logrus.Error("update zhihu hot_info err:", err)
					//fmt.Printf("Failed to set value in Redis: %v", err)
					return

				}
			}
		}
	}
}
func Refresh() []ZhiHu {
	var maxUpdateVer int64

	// 查询最大的 update_ver
	result := model.Conn.Model(&ZhiHu{}).Select("MAX(update_ver) as max_update_ver").Scan(&maxUpdateVer)
	if result.Error != nil {
		logrus.Error("zhihu:Refresh:1:", result.Error)
		//log.Fatal(result.Error)
	}

	// 查询所有 update_ver 为最大值的记录
	var zhihuList []ZhiHu
	result = model.Conn.Where("update_ver = ?", maxUpdateVer).Find(&zhihuList)
	if result.Error != nil {
		logrus.Error("zhihu:Refresh:2:", result.Error)
		//log.Fatal(result.Error)
	}

	// 打印查询结果
	fmt.Printf("Data with max update_ver (%d):\n", maxUpdateVer)
	return zhihuList

}
