package toutiao

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"hotinfo/app/model"
	"hotinfo/app/tools"
	"io/ioutil"
	"net/http"
	"time"
)

func init() {
	// 设置配置文件名和路径
	viper.SetConfigName("config.yaml") // 配置文件名（不含扩展名）
	viper.SetConfigType("yaml")        // 配置文件类型
	viper.AddConfigPath(".")           // 配置文件所在路径

	err := viper.ReadInConfig()
	if err != nil {
		panic("配置文件读取失败")
	}
}
func Run() {
	ticker := time.NewTicker(10 * time.Minute)
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
	request, err := http.NewRequest("GET", viper.GetString("hot_api.toutiao"), nil)
	if err != nil {
		logrus.Error("toutiao:Failed to create HTTP request:", err)
		//fmt.Println("Failed to create HTTP request:", err)
		return
	}

	// 发送请求并获取响应
	response, err := client.Do(request)
	if err != nil {
		logrus.Error("toutiao:HTTP request failed:", err)
		//fmt.Println("HTTP request failed:", err)
		return
	}
	defer response.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logrus.Error("toutiao:Failed to read response body:", err)
		//fmt.Println("Failed to read response body:", err)
		return
	}

	// 解析响应内容到结构体
	var baiduHot Hot
	err = json.Unmarshal(body, &baiduHot)
	if err != nil {
		logrus.Error("toutiao:Failed to unmarshal response body:", err)
		//fmt.Println("Failed to unmarshal response body:", err)
		return
	}
	var data []TouTiao
	now := time.Now().Unix()
	var hotinfoStr string
	for _, datum := range baiduHot.List {
		a := TouTiao{
			Icon:        datum.Label,
			LabelDesc:   datum.LabelDesc,
			UpdateVer:   now,
			Title:       datum.Title,
			Url:         datum.Url,
			Hot:         datum.Hot,
			CreatedTime: time.Now(),
			UpdatedTime: time.Now(),
		}
		data = append(data, a)
		hotinfoStr = hotinfoStr + datum.Title + datum.Url
	}
	hashStr := tools.Sha256Hash(hotinfoStr)

	value, err := model.RedisClient.Get(context.Background(), "toutiao_hot").Result()
	if err == redis.Nil {
		err = model.RedisClient.Set(context.Background(), "toutiao_hot", hashStr, 0).Err()
		if err != nil {
			logrus.Error("toutiao:Failed to set value in Redis:", err)
			//fmt.Printf("Failed to set value in Redis: %v", err)
			return

		}
		model.Conn.Create(data)
	} else if err != nil {
		logrus.Error("toutiao:Error getting value from Redis:", err)
		//fmt.Printf("Error getting value from Redis: %v", err)
	} else {

		if hashStr != value {
			err = model.RedisClient.Set(context.Background(), "toutiao_hot", hashStr, 0).Err()
			if err != nil {
				logrus.Error("toutiao:Error setting value from Redis:", err)
			}
			err = model.Conn.Create(data).Error
			if err != nil {
				logrus.Error("toutiao:db_create:", err)
			}
		} else {
			var maxUpdateVer int64
			var updateSlice []TouTiao
			model.Conn.Model(&TouTiao{}).Select("MAX(update_ver) as max_update_ver").Scan(&maxUpdateVer)
			model.Conn.Where("update_ver = ?", maxUpdateVer).Find(updateSlice)
			for _, record := range updateSlice {
				record.UpdateVer = now
				record.UpdatedTime = time.Now()
				err := model.Conn.Save(&record).Error
				if err != nil {
					logrus.Error("update toutiao hot_info err:", err)
					//fmt.Printf("Failed to set value in Redis: %v", err)
					return

				}
			}
		}
	}

}
func Refresh() []TouTiao {
	var maxUpdateVer int64

	// 查询最大的 update_ver
	result := model.Conn.Model(&TouTiao{}).Select("MAX(update_ver) as max_update_ver").Scan(&maxUpdateVer)
	if result.Error != nil {
		logrus.Error("toutiao:Refresh:1:", result.Error)
		//log.Fatal(result.Error)
	}

	// 查询所有 update_ver 为最大值的记录
	var toutiaoList []TouTiao
	result = model.Conn.Where("update_ver = ?", maxUpdateVer).Find(&toutiaoList)
	if result.Error != nil {
		logrus.Error("toutiao:Refresh:2:", result.Error)
		//log.Fatal(result.Error)
	}

	// 打印查询结果
	fmt.Printf("Data with max update_ver (%d):\n", maxUpdateVer)
	return toutiaoList

}
