package juejin

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"hotinfo/app/model"
	"hotinfo/app/tools"
	"io"
	"net/http"
	"time"
)

const api = "https://api.juejin.cn/content_api/v1/content/article_rank?category_id=1&type=hot&aid=2608&uuid=7263387509600912950&spider=0"

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
	client := &http.Client{}
	request, err := http.NewRequest("GET", api, nil)
	if err != nil {
		logrus.Error("juejin:Error creating request:", err)
		//fmt.Println("Error creating request:", err)
		return
	}
	request.Header.Add("User-Agent", "Not_A Brand\";v=\"8\", \"Chromium\";v=\"120\", \"Microsoft Edge\";v=\"120")
	request.Header.Add("Cookie", "_tea_utm_cache_2608=undefined; __tea_cookie_tokens_2608=%257B%2522web_id%2522%253A%25227263387509600912950%2522%252C%2522user_unique_id%2522%253A%25227263387509600912950%2522%252C%2522timestamp%2522%253A1691139204499%257D; uid_tt=d9129948ced9f65f755090d09acb09b9; uid_tt_ss=d9129948ced9f65f755090d09acb09b9; sid_tt=410d35bde6e42e93cc82a7b7d86cad0b; sessionid=410d35bde6e42e93cc82a7b7d86cad0b; sessionid_ss=410d35bde6e42e93cc82a7b7d86cad0b; sid_guard=410d35bde6e42e93cc82a7b7d86cad0b%7C1693644523%7C31535999%7CSun%2C+01-Sep-2024+08%3A48%3A42+GMT; sid_ucp_v1=1.0.0-KGE2ZjRmMDYwNWI5ODE4Yzk1ZWUyYzgxZGQ2NDY3ZGNmMzY5YWI1NDMKFgjoq_CLrc3gAhDr7cunBhiwFDgIQAsaAmxmIiA0MTBkMzViZGU2ZTQyZTkzY2M4MmE3YjdkODZjYWQwYg; ssid_ucp_v1=1.0.0-KGE2ZjRmMDYwNWI5ODE4Yzk1ZWUyYzgxZGQ2NDY3ZGNmMzY5YWI1NDMKFgjoq_CLrc3gAhDr7cunBhiwFDgIQAsaAmxmIiA0MTBkMzViZGU2ZTQyZTkzY2M4MmE3YjdkODZjYWQwYg; store-region=cn-ha; store-region-src=uid; csrf_session_id=a1ad739a0df8eb5fe9da982694aff9be; msToken=6hijmTqCkBVaeUeDp-YNfKGrJFKsNtkCIQFytMyj20fAWZcaovIZgjh1SKrttjR8pBCQh_c8rARGVcDzscbzlqNoThNENwjMVMuD6Rt3gY7-2vGrTx20Uqb4ZVLNimqr")
	resp, err := client.Do(request)
	if err != nil {
		logrus.Error("juejin:请求失败:", err)
		//fmt.Println("请求失败:", err)
		return
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Error("juejin:Error reading response body:", err)
		//fmt.Println("Error reading response body:", err)
		return
	}
	//fmt.Printf("body:%s", string(body))

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		logrus.Error("juejin:Error parsing JSON:", err)
		//fmt.Println("Error parsing JSON:", err)
		return
	}

	// 获取所有的 realtime 数据
	realtimeData := response.Data

	// 打印解析后的 realtime 数据
	//fmt.Printf("Realtime data: %+v\n", realtimeData)
	data := make([]*Juejin, 0)
	now := time.Now().Unix()
	head := "https://juejin.cn/post/"
	var hotinfoStr string
	for _, list := range realtimeData {
		url := head + list.Content.ContentID
		tmp := Juejin{
			UpdateVer:   now,
			Hot:         list.ContentCounter.HotRank,
			Title:       list.Content.Title,
			Url:         url,
			AuthorName:  list.Author.Name,
			AuthorId:    list.Author.UserID,
			CreatedTime: time.Now(),
			UpdatedTime: time.Now(),
		}
		data = append(data, &tmp)
		hotinfoStr = list.Content.Title + url + hotinfoStr
	}
	hashStr := tools.Sha256Hash(hotinfoStr)

	value, err := model.RedisClient.Get(context.Background(), "juejin_hot").Result()
	if err == redis.Nil {
		err = model.RedisClient.Set(context.Background(), "juejin_hot", hashStr, 0).Err()
		if err != nil {
			logrus.Error("juejin:Failed to set value in Redis:", err)
			//fmt.Printf("Failed to set value in Redis: %v", err)
			return

		}
		model.Conn.Create(data)
	} else if err != nil {
		logrus.Error("juejin:Error getting value from Redis:", err)
		//fmt.Printf("Error getting value from Redis: %v", err)
	} else {

		if hashStr != value {
			err = model.RedisClient.Set(context.Background(), "juejin_hot", hashStr, 0).Err()
			if err != nil {
				logrus.Error("juejin:Error setting value from Redis:", err)
			}
			err = model.Conn.Create(data).Error
			if err != nil {
				logrus.Error("juejin:db_create:", err)
			}
		} else {
			var maxUpdateVer int64
			var updateSlice []Juejin
			model.Conn.Model(&Juejin{}).Select("MAX(update_ver) as max_update_ver").Scan(&maxUpdateVer)
			model.Conn.Where("update_ver = ?", maxUpdateVer).Find(updateSlice)
			for _, record := range updateSlice {
				record.UpdateVer = now
				record.UpdatedTime = time.Now()
				err := model.Conn.Save(&record).Error
				if err != nil {
					logrus.Error("update juejin hot_info err:", err)
					//fmt.Printf("Failed to set value in Redis: %v", err)
					return

				}
			}
		}
	}

}
func Refresh() []Juejin {
	var maxUpdateVer int64

	// 查询最大的 update_ver
	result := model.Conn.Model(&Juejin{}).Select("MAX(update_ver) as max_update_ver").Scan(&maxUpdateVer)
	if result.Error != nil {
		logrus.Error("juejin:Refresh:1:", result.Error)
		//log.Fatal(result.Error)
	}

	// 查询所有 update_ver 为最大值的记录
	var juejinList []Juejin
	result = model.Conn.Where("update_ver = ?", maxUpdateVer).Find(&juejinList)
	if result.Error != nil {
		logrus.Error("juejin:Refresh:2:", result.Error)
		//log.Fatal(result.Error)
	}

	// 打印查询结果
	fmt.Printf("Data with max update_ver (%d):\n", maxUpdateVer)
	return juejinList

}
