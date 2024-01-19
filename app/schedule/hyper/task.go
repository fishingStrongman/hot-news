package hyper

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"hotinfo/app/model"
	"hotinfo/app/tools"
	"io"
	"log"
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

func getFirstTag(tagList []*TagList) string {
	if len(tagList) > 0 {
		return tagList[0].Tag
	}
	return ""
}

func getInfo() {
	client := &http.Client{}
	request, err := http.NewRequest("GET", viper.GetString("hot_api.hyper"), nil)
	if err != nil {
		logrus.Error("hyper:Error creating request:", err)
		//fmt.Println("hyper:Error creating request:", err)
		return
	}

	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	request.Header.Add("Cookie", "Hm_lvt_94a1e06bbce219d29285cee2e37d1d26=1704782716,1704888210; ariaDefaultTheme=undefined; Hm_lpvt_94a1e06bbce219d29285cee2e37d1d26=1704888327")
	response, err := client.Do(request)
	if err != nil {
		logrus.Error("hyper:Error sending request:", err)
		//fmt.Println("Error sending request:", err)
		return
	}
	defer func() {
		_ = response.Body.Close()
	}()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		logrus.Error("hyper:Error reading response body:", err)
		//fmt.Println("Error reading response body:", err)
		return
	}
	//fmt.Printf("body:%s", string(body))
	var ret Ret
	_ = json.Unmarshal(body, &ret)
	//fmt.Printf("ret:%v", ret.Data.HotNews)
	data := make([]*Hyper, 0)
	now := time.Now().Unix()
	var hotinfoStr string
	for _, result := range ret.Data.HotNews {
		url := "https://www.thepaper.cn/newsDetail_forward_" + result.ContId
		tmp := Hyper{
			UpdateVer:   now,
			Title:       result.Name,
			KeyWord:     getFirstTag(result.TagList),
			Url:         url,
			CreatedTime: time.Now(),
			UpdatedTime: time.Now(),
		}
		data = append(data, &tmp)
		hotinfoStr = result.Name + url + hotinfoStr
	}
	hashStr := tools.Sha256Hash(hotinfoStr)

	value, err := model.RedisClient.Get(context.Background(), "hyper_hot").Result()
	if err == redis.Nil {
		err = model.RedisClient.Set(context.Background(), "hyper_hot", hashStr, 0).Err()
		if err != nil {
			logrus.Error("hyper:Failed to set value in Redis:", err)
			//fmt.Printf("Failed to set value in Redis: %v", err)
			return

		}
		model.Conn.Create(data)
	} else if err != nil {
		logrus.Error("hyper:Error getting value from Redis:", err)
		//fmt.Printf("Error getting value from Redis: %v", err)
	} else {

		if hashStr != value {
			err = model.RedisClient.Set(context.Background(), "hyper_hot", hashStr, 0).Err()
			if err != nil {
				logrus.Error("hyper:Error setting value from Redis:", err)
			}
			err = model.Conn.Create(data).Error
			if err != nil {
				logrus.Error("hyper:db_create:", err)
			}
		} else {
			var maxUpdateVer int64
			var updateSlice []Hyper
			model.Conn.Model(&Hyper{}).Select("MAX(update_ver) as max_update_ver").Scan(&maxUpdateVer)
			model.Conn.Where("update_ver = ?", maxUpdateVer).Find(&updateSlice)
			for _, record := range updateSlice {
				record.UpdateVer = now
				record.UpdatedTime = time.Now()
				err := model.Conn.Save(&record).Error
				if err != nil {
					logrus.Error("update hyper hot_info err:", err)
					//fmt.Printf("Failed to set value in Redis: %v", err)
					return

				}
			}
		}
	}
}
func Refresh() []Hyper {
	var maxUpdateVer int64

	// 查询最大的 update_ver
	result := model.Conn.Model(&Hyper{}).Select("MAX(update_ver) as max_update_ver").Scan(&maxUpdateVer)
	if result.Error != nil {
		logrus.Error("hyper:Refresh:1:", result.Error)
		log.Fatal(result.Error)
	}

	// 查询所有 update_ver 为最大值的记录
	var pengpaiList []Hyper
	result = model.Conn.Where("update_ver = ?", maxUpdateVer).Find(&pengpaiList)
	if result.Error != nil {
		logrus.Error("hyper:Refresh:2:", result.Error)
		//log.Fatal(result.Error)
	}

	// 打印查询结果
	fmt.Printf("Data with max update_ver (%d):\n", maxUpdateVer)
	return pengpaiList

}
