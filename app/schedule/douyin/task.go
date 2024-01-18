package douyin

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

const api = "https://www.douyin.com/aweme/v1/web/hot/search/list/?device_platform=webapp&aid=6383&channel=channel_pc_web&detail_list=1&source=6&board_type=0&board_sub_type=&pc_client_type=1&version_code=170400&version_name=17.4.0&cookie_enabled=true&screen_width=1536&screen_height=864&browser_language=zh-CN&browser_platform=Win32&browser_name=Edge&browser_version=116.0.1938.76&browser_online=true&engine_name=Blink&engine_version=116.0.0.0&os_name=Windows&os_version=10&cpu_core_num=16&device_memory=8&platform=PC&downlink=10&effective_type=4g&round_trip_time=200&webid=7321996960310937088&msToken=2PE2cwtw2KA_3xc3c1KxDTbkVwCPvWZRXh2Oik0TnjElG-Fn0VvHeFjycKRPRHQ6p71nVgIHEHkEE3pkaaY9t22pOLglxxmnyeH20H11_1rOzY9IuQ==&X-Bogus=DFSzswVLvtUANyuHt7En9ENSwbuS"

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
		logrus.Error("douyin:Failed to create HTTP request:", err)
		//fmt.Println("Failed to create HTTP request:", err)
		return
	}
	//添加User-Agent
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36 Edg/116.0.1938.76")
	//添加Cookie
	request.Header.Add("Cookie", "ttwid=1%7CZfTDRjzLikJLdXf5vUt-PHdoxm50XhdPorB3mEo4lpw%7C1704785277%7Cce482bdd0bd65873b36caf586eb7c9045d7833732ee7f7dd3b058b1ef28e1422; douyin.com; device_web_cpu_core=16; device_web_memory_size=8; architecture=amd64; dy_swidth=1536; dy_sheight=864; csrf_session_id=a1ad739a0df8eb5fe9da982694aff9be; strategyABtestKey=%221704785278.681%22; s_v_web_id=verify_lr614zum_BIQ8Wfu6_jgYj_4uHu_8mXd_XqCjU8M0yyU6; volume_info=%7B%22isUserMute%22%3Afalse%2C%22isMute%22%3Atrue%2C%22volume%22%3A0.5%7D; passport_csrf_token=f12476d310a7abece48ccf2000fb1191; passport_csrf_token_default=f12476d310a7abece48ccf2000fb1191; bd_ticket_guard_client_web_domain=2; ttcid=7b03da3762684406893a3f32208385dc41; bd_ticket_guard_client_data=eyJiZC10aWNrZXQtZ3VhcmQtdmVyc2lvbiI6MiwiYmQtdGlja2V0LWd1YXJkLWl0ZXJhdGlvbi12ZXJzaW9uIjoxLCJiZC10aWNrZXQtZ3VhcmQtcmVlLXB1YmxpYy1rZXkiOiJCS0VVT3ZCMkNSMnp5NTFXMUIvd1BnbkpSQWFKTGpzbCtkdC9hN2I4eTNZOVhaT0M3M3R0MW9nWXBKU0hNSDhkMkJBV3BvUU4xYWgzblJLS2gyQkhZazA9IiwiYmQtdGlja2V0LWd1YXJkLXdlYi12ZXJzaW9uIjoxfQ%3D%3D; FORCE_LOGIN=%7B%22videoConsumedRemainSeconds%22%3A180%2C%22isForcePopClose%22%3A1%7D; SEARCH_RESULT_LIST_TYPE=%22single%22; stream_player_status_params=%22%7B%5C%22is_auto_play%5C%22%3A0%2C%5C%22is_full_screen%5C%22%3A0%2C%5C%22is_full_webscreen%5C%22%3A0%2C%5C%22is_mute%5C%22%3A1%2C%5C%22is_speed%5C%22%3A1%2C%5C%22is_visible%5C%22%3A1%7D%22; download_guide=%223%2F20240109%2F0%22; pwa2=%220%7C0%7C2%7C0%22; __ac_nonce=0659d291c005f5a1850c3; __ac_signature=_02B4Z6wo00f01JHQKBgAAIDDxFKjh9DXr7yR8CyAAEHtp4exxEb809p7RSpC.w.F3L0FHh30yMxQho9R68PekYFwakelqedtamPdgojEStl5o24Ja5j2mnVJBXHzpGqbydNfmJi8xYHEUfwL1c; IsDouyinActive=true; stream_recommend_feed_params=%22%7B%5C%22cookie_enabled%5C%22%3Atrue%2C%5C%22screen_width%5C%22%3A1536%2C%5C%22screen_height%5C%22%3A864%2C%5C%22browser_online%5C%22%3Atrue%2C%5C%22cpu_core_num%5C%22%3A16%2C%5C%22device_memory%5C%22%3A8%2C%5C%22downlink%5C%22%3A10%2C%5C%22effective_type%5C%22%3A%5C%224g%5C%22%2C%5C%22round_trip_time%5C%22%3A50%7D%22; home_can_add_dy_2_desktop=%221%22; msToken=FfdOKnbchGkVoZiQZD7kkuyyOaIPoQi-REGWK_iPkQ6M0FReZ_U8jpC0znarZdetvuwm4ci7scr6wJlhMTcKduXusF6UULypEWL02etM8wDJJ_GIRQ==; msToken=qtTgVJDncrAxogJwklf8JK0ZteGMiSNmq-pNfhtJWdU9Mvz4JvELg5pCqGdlptpQEsYsAnUW2mj12EweyhP62eXkX6OhX9Gls3J0T-XBVfORVSo0Yw==; tt_scid=SVzYkwrJ9v54XP85PprplUQdaTjaTGzsSloYk5-SNwbZvahTdKAS5ThgWFqBNcJ903f4")
	// 发送请求并获取响应
	response, err := client.Do(request)
	if err != nil {
		logrus.Error("douyin:HTTP request failed:", err)
		//fmt.Println("HTTP request failed:", err)
		return
	}
	defer response.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logrus.Error("douyin:Failed to read response body:", err)
		//fmt.Println("Failed to read response body:", err)
		return
	}
	// 解析响应内容到结构体
	var baiduHot Hot
	err = json.Unmarshal(body, &baiduHot)
	if err != nil {
		logrus.Error("douyin:Failed to unmarshal response body:", err)
		//fmt.Println("Failed to unmarshal response body:", err)
		return
	}
	var data []DouYin
	now := time.Now().Unix()
	var hotinfoStr string
	for _, datum := range baiduHot.Data.List {
		a := DouYin{
			UpdateVer:   now,
			Title:       datum.Title,
			Url:         "https://www.douyin.com/search/" + datum.Title,
			Hot:         datum.Hot,
			CreatedTime: time.Now(),
			UpdatedTime: time.Now(),
		}
		data = append(data, a)
		hotinfoStr = datum.Title + "https://www.douyin.com/search/" + datum.Title + hotinfoStr
	}
	hashStr := tools.Sha256Hash(hotinfoStr)

	value, err := model.RedisClient.Get(context.Background(), "douyin_hot").Result()
	if err == redis.Nil {
		err = model.RedisClient.Set(context.Background(), "douyin_hot", hashStr, 0).Err()
		if err != nil {
			logrus.Error("douyin:Failed to set value in Redis:", err)
			//fmt.Printf("Failed to set value in Redis: %v", err)
			return

		}
		model.Conn.Create(data)
	} else if err != nil {
		logrus.Error("douyin:Error getting value from Redis:", err)
		//fmt.Printf("Error getting value from Redis: %v", err)
	} else {

		if hashStr != value {
			err = model.RedisClient.Set(context.Background(), "douyin_hot", hashStr, 0).Err()
			if err != nil {
				logrus.Error("douyin:Error setting value from Redis:", err)
			}
			err = model.Conn.Create(data).Error
			if err != nil {
				logrus.Error("douyin:db_create:", err)
			}
		} else {
			var maxUpdateVer int64
			var updateSlice []DouYin
			model.Conn.Model(&DouYin{}).Select("MAX(update_ver) as max_update_ver").Scan(&maxUpdateVer)
			model.Conn.Where("update_ver = ?", maxUpdateVer).Find(&updateSlice)
			for _, record := range updateSlice {
				record.UpdateVer = now
				record.UpdatedTime = time.Now()
				err := model.Conn.Save(&record).Error
				if err != nil {
					logrus.Error("update douyin hot_info err:", err)
					//fmt.Printf("Failed to set value in Redis: %v", err)
					return

				}
			}
		}
	}

}
func Refresh() []DouYin {
	var maxUpdateVer int64

	// 查询最大的 update_ver
	result := model.Conn.Model(&DouYin{}).Select("MAX(update_ver) as max_update_ver").Scan(&maxUpdateVer)
	if result.Error != nil {
		logrus.Error("douyin:Refresh:1:", result.Error)
		//log.Fatal(result.Error)
	}

	// 查询所有 update_ver 为最大值的记录
	var douyinList []DouYin
	result = model.Conn.Where("update_ver = ?", maxUpdateVer).Find(&douyinList)
	if result.Error != nil {
		logrus.Error("douyin:Refresh:2:", result.Error)
		//log.Fatal(result.Error)
	}

	// 打印查询结果
	fmt.Printf("Data with max update_ver (%d):\n", maxUpdateVer)
	return douyinList

}
