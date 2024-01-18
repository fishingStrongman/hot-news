package bilibili_rank

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

// const api = "https://api.bilibili.com/x/web-interface/ranking/v2?rid=0&type=all&web_location=333.934&w_rid=a6a74f54c01bf69f693d80f2da05e329&wts=1704943132"
const api = "https://api.bilibili.com/x/web-interface/ranking/v2"

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
		logrus.Error("bilibili_rank:Error creating request:", err)
		return
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	request.Header.Add("Cookie", "buvid3=43094C0F-3E29-EEEC-7AF9-71709B5FB1D330870infoc; b_nut=1689662630; i-wanna-go-back=-1; _uuid=258B7F9B-3FAB-5833-EF4D-E34410EFCFA6136563infoc; FEED_LIVE_VERSION=V8; rpdid=|(~ummuR|k~0J'uY)mYYkY~l; DedeUserID=39362173; DedeUserID__ckMd5=0a62ae35f67fb3ec; buvid4=DFD3FA46-D7A5-1705-5B4D-6FA923F5A8BA45558-023071814-hTRASHGPQlzJ6%2BUUl9tIRw%3D%3D; header_theme_version=CLOSE; buvid_fp_plain=undefined; hit-new-style-dyn=1; hit-dyn-v2=1; b_ut=5; is-2022-channel=1; enable_web_push=DISABLE; CURRENT_BLACKGAP=0; CURRENT_FNVAL=4048; dy_spec_agreed=1; CURRENT_QUALITY=80; fingerprint=4d23e39272c183e66fde0d7793bc93a2; buvid_fp=4f9bfd389e80ed5e3b4daa8231fb596f; LIVE_BUVID=AUTO3817037443317552; bili_ticket=eyJhbGciOiJIUzI1NiIsImtpZCI6InMwMyIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDQ5NjEzOTMsImlhdCI6MTcwNDcwMjEzMywicGx0IjotMX0.5jYWcWYlWdJAt8ubGPESTrb-cxL5AnzNgK83cKqR3As; bili_ticket_expires=1704961333; home_feed_column=5; browser_resolution=1920-969; SESSDATA=f28f92ab%2C1720487121%2Cbb06a%2A12CjDbycLLcqhZGyxdoxgLWPAh5a3jz4v6EiAGEbyaO9IV3fnzj5YJQOE7t4yur1IAzLUSVjg3Zjh5bHVaeEo3SGtpNVdfdVhiTjBscFRKMDZjcVdEZUpuSHkzREg1MzRCcERmanJNMzB4eVhYQmhQTHAwSVh3N0lMLWFKOVRFVGR5TnRIZloyR3J3IIEC; bili_jct=3fed309212edca07cd3057d23ee85d2a; sid=7913l753; bp_video_offset_39362173=884953491410255958; PVID=4; b_lsid=7A2C107BE_18CF6A671CE; innersign=0")
	response, err := client.Do(request)
	if err != nil {
		logrus.Error("bilibili_rank:Error sending request:", err)
		return
	}
	defer func() {
		_ = response.Body.Close()
	}()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		logrus.Error("bilibili_rank:Error reading response body:", err)
		return
	}
	//fmt.Printf("body:%s", string(body))
	var ret Ret
	_ = json.Unmarshal(body, &ret)
	data := make([]*BRank, 0)
	now := time.Now().Unix()
	var hotinfoStr string
	for _, result := range ret.Data.List {
		tmp := BRank{
			UpdateVer:   now,
			Title:       result.Title,
			Tag:         result.TName,
			Author:      result.Owner.Name,
			Url:         result.ShortLinkV2,
			CreatedTime: time.Now(),
			UpdatedTime: time.Now(),
		}
		data = append(data, &tmp)
		hotinfoStr = result.Title + result.TName + hotinfoStr
	}
	hashStr := tools.Sha256Hash(hotinfoStr)

	value, err := model.RedisClient.Get(context.Background(), "bilibili_rank_hot").Result()
	if err == redis.Nil {
		err = model.RedisClient.Set(context.Background(), "bilibili_rank_hot", hashStr, 0).Err()
		if err != nil {
			logrus.Error("bilibili_rank:Failed to set value in Redis:", err)
			//fmt.Printf("Failed to set value in Redis: %v", err)
			return

		}
		model.Conn.Create(data)
	} else if err != nil {
		logrus.Error("bilibili_rank:Error getting value from Redis:", err)
		//fmt.Printf("Error getting value from Redis: %v", err)
	} else {

		if hashStr != value {
			err = model.RedisClient.Set(context.Background(), "bilibili_rank_hot", hashStr, 0).Err()
			if err != nil {
				logrus.Error("bilibili_rank_hot:Error setting value from Redis:", err)
			}
			err = model.Conn.Create(data).Error
			if err != nil {
				logrus.Error("bilibili_rank_hot:db_create:", err)
			}
		} else {
			var maxUpdateVer int64
			var updateSlice []BRank
			model.Conn.Model(&BRank{}).Select("MAX(update_ver) as max_update_ver").Scan(&maxUpdateVer)
			model.Conn.Where("update_ver = ?", maxUpdateVer).Find(updateSlice)
			for _, record := range updateSlice {
				record.UpdateVer = now
				record.UpdatedTime = time.Now()
				err := model.Conn.Save(&record).Error
				if err != nil {
					logrus.Error("update bilibili_rank hot_info err:", err)
					//fmt.Printf("Failed to set value in Redis: %v", err)
					return

				}
			}
		}
	}

}
func Refresh() []BRank {
	var maxUpdateVer int64

	// 查询最大的 update_ver
	result := model.Conn.Model(&BRank{}).Select("MAX(update_ver) as max_update_ver").Scan(&maxUpdateVer)
	if result.Error != nil {
		logrus.Error("bilibili_rank:Refresh:1:", result.Error)
	}

	// 查询所有 update_ver 为最大值的记录
	var bilibiliRankList []BRank
	result = model.Conn.Where("update_ver = ?", maxUpdateVer).Find(&bilibiliRankList)
	if result.Error != nil {
		logrus.Error("bilibili_rank:Refresh:2:", result.Error)
	}

	// 打印查询结果
	fmt.Printf("Data with max update_ver (%d):\n", maxUpdateVer)
	return bilibiliRankList

}
