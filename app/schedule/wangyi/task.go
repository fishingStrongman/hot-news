package wangyi

import (
	"encoding/json"
	"fmt"
	"hotinfo/app/model"
	"io"
	"log"
	"net/http"
	"time"
)

const api = "https://m.163.com/fe/api/hot/news/flow"

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
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36 Edg/120.0.0.0")
	request.Header.Add("Cookie", "Hm_lvt_b2d0b085a122275dd543c6d39d92bc62=1697247591; Hm_lvt_1e8bfbf2a9c357adc57abdfb1fddd6f8=1697247592; Hm_lvt_ee16af73ee73a16e2cf07eba7a4f152b=1697247592; Hm_lvt_80e4acbac28178f85c9da26879bb2070=1697247592; Hm_lvt_c90426000c3f454ef4c16be54e22ae34=1697247592; NTES_YD_PASSPORT=od5sfqcTfBOXNIWnGsGjPSypYiPGdm9vS4OPTrrDq6lTbX4tbOxW2S0TGjqz29d7OkGEqkWgifxBbNaTQMp.fVSe.T4JZhSXdZRAIkOx5HKk1U3Uq3awIUShUPBfpf5lPl1kFFNi4J6XObLCau0nbDJ7RgO148LeGOhf9WBGjMXN6aqizEfClCrtmhVml_WOlKxUj8xatLtjDhrBVsI1lcQKL; P_INFO=13783132602|1704805855|0|163|00&99|null&null&null#hen&410700#10#0#0|&0||13783132602")
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("请求失败:", err)
		return
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	//fmt.Printf("body:%s", string(body))

	var response T
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// 获取所有的 realtime 数据
	realtimeData := response.Data.List

	// 打印解析后的 realtime 数据
	//fmt.Printf("Realtime data: %+v\n", realtimeData)
	data := make([]*WangYi, 0)
	now := time.Now().Unix()

	for _, list := range realtimeData {

		tmp := WangYi{
			UpdateVer:   now,
			Title:       list.Title,
			Url:         list.Url,
			KeyWord:     list.Keyword,
			CreatedTime: time.Now(),
			UpdatedTime: time.Now(),
		}
		data = append(data, &tmp)

	}
	model.Conn.Create(data)
}
func Refresh() []WangYi {
	var maxUpdateVer int64

	// 查询最大的 update_ver
	result := model.Conn.Model(&WangYi{}).Select("MAX(update_ver) as max_update_ver").Scan(&maxUpdateVer)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	// 查询所有 update_ver 为最大值的记录
	var wangyiList []WangYi
	result = model.Conn.Where("update_ver = ?", maxUpdateVer).Find(&wangyiList)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	// 打印查询结果
	fmt.Printf("Data with max update_ver (%d):\n", maxUpdateVer)
	return wangyiList

}
