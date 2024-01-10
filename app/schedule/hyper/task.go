package hyper

import (
	"encoding/json"
	"fmt"
	"hotinfo/app/model"
	"io"
	"net/http"
	"time"
)

const api = "https://cache.thepaper.cn/contentapi/wwwIndex/rightSidebar"

func Run() {
	getHyper()
}
func getFirstTag(tagList []*TagList) string {
	if len(tagList) > 0 {
		return tagList[0].Tag
	}
	return ""
}
func getHyper() {
	client := &http.Client{}
	request, err := http.NewRequest("GET", api, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	request.Header.Add("Cookie", "Hm_lvt_94a1e06bbce219d29285cee2e37d1d26=1704782716,1704888210; ariaDefaultTheme=undefined; Hm_lpvt_94a1e06bbce219d29285cee2e37d1d26=1704888327")
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
	//fmt.Printf("body:%s", string(body))
	var ret Ret
	_ = json.Unmarshal(body, &ret)
	//fmt.Printf("ret:%v", ret.Data.HotNews)
	data := make([]*Hyper, 0)
	for _, result := range ret.Data.HotNews {
		url := "https://www.thepaper.cn/newsDetail_forward_" + result.ContId
		tmp := Hyper{
			UpdateVer:   time.Now().Unix(),
			Title:       result.Name,
			KeyWord:     getFirstTag(result.TagList),
			Url:         url,
			CreatedTime: time.Now(),
			UpdatedTime: time.Now(),
		}
		data = append(data, &tmp)
	}
	model.Conn.Create(data)
}
