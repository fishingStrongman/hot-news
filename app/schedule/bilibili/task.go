package bilibili

import (
	"encoding/json"
	"fmt"
	"hotinfo/app/model"
	"io"
	"log"
	"net/http"
	"time"
)

//api: https://api.bilibili.com/x/web-interface/wbi/search/square?limit=20&platform=web&w_rid=18e2eb4d6b330093545209bee0f28408&wts=1704525897

const api = "https://api.bilibili.com/x/web-interface/wbi/search/square?limit=20&platform=web"

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
	client := &http.Client{}
	request, err := http.NewRequest("GET", api, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	defer func() {
		_ = response.Body.Close()
	}()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Printf("body:%s", string(body))

	//body, _ := os.ReadFile("./app/schedule/bilibili/data.json")
	//fmt.Printf("body:%s\n", body)

	var ret Ret
	_ = json.Unmarshal(body, &ret)
	fmt.Printf("ret:%v", ret.Data.Trending)
	data := make([]*Bilibili, 0)
	now := time.Now().Unix()
	for _, list := range ret.Data.Trending.List {
		fmt.Printf("list:%+v\n", list)
		tmp := Bilibili{
			UpdateVer:   now,
			Title:       list.ShowName,
			Icon:        list.Icon,
			KeyWord:     list.Keyword,
			CreatedTime: time.Now(),
			UpdatedTime: time.Now(),
		}
		data = append(data, &tmp)
	}

	model.Conn.Create(data)
}
func Refresh() []Bilibili {
	var maxUpdateVer int64

	// 查询最大的 update_ver
	result := model.Conn.Model(&Bilibili{}).Select("MAX(update_ver) as max_update_ver").Scan(&maxUpdateVer)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	// 查询所有 update_ver 为最大值的记录
	var bilibiliList []Bilibili
	result = model.Conn.Where("update_ver = ?", maxUpdateVer).Find(&bilibiliList)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	// 打印查询结果
	fmt.Printf("Data with max update_ver (%d):\n", maxUpdateVer)
	return bilibiliList

}
