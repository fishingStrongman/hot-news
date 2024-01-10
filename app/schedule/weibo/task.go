package weibo

import (
	"encoding/json"
	"fmt"
	"hotinfo/app/model"
	"io"
	"net/http"
	"time"
)

const api = "https://weibo.com/ajax/side/hotSearch"

func Run() {
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
	request.Header.Add("Cookie", "SINAGLOBAL=5477913124167.391.1697440356322; SUB=_2AkMSx9VZf8NxqwFRmfoVzmjjbYp1zQ_EieKkmySCJRMxHRl-yT9vqkgstRB6OUf7tcleCDNJ6arIETm4cYrHXw_0sMyk; SUBP=0033WrSXqPxfM72-Ws9jqgMF55529P9D9WW28SUoBDEHpfe5GOIDOyWD; UOR=www.fengfengzhidao.com,s.weibo.com,cn.bing.com; WBPSESS=V0zdZ7jH8_6F0CA8c_ussawu4Ddto8Lqji8oSVM-Jdeek5bWx1mzySd1GR2Gdf26WG8Q_ihMCaE90eNYOZ1IZ3P2T5qVc2jvoggT0M8k30aEtBZXpnvP13rMuk4UkEAZ92yvZ_jTiFuaSVy00oieKtft309-NVOQs7WlJyItUwI=; ULV=1704762642630:6:3:3:3031869891888.5034.1704762642575:1704701239233; XSRF-TOKEN=0Cc5Yz_3DBT2jsAc0iuxANGy")
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
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// 获取所有的 realtime 数据
	realtimeData := response.Data.Realtime

	// 打印解析后的 realtime 数据
	fmt.Printf("Realtime data: %+v\n", realtimeData)
	data := make([]*WeiBo, 0)
	now := time.Now().Unix()
	head := "https://s.weibo.com/weibo?q=%23"
	tail := "%23&t=31&band_rank=2&Refer=top"
	for _, list := range realtimeData {
		url := head + list.Note + tail
		fmt.Println(url)
		tmp := WeiBo{
			UpdateVer:   now,
			Note:        list.Note,
			Url:         url,
			IconDesc:    list.IconDesc,
			Category:    list.Category,
			CreatedTime: time.Now(),
			UpdatedTime: time.Now(),
		}
		data = append(data, &tmp)

	}
	model.Conn.Create(data)

}
