package juejin

import (
	"encoding/json"
	"fmt"
	"hotinfo/app/model"
	"io"
	"net/http"
	"time"
)

const api = "https://api.juejin.cn/content_api/v1/content/article_rank?category_id=1&type=hot&aid=2608&uuid=7263387509600912950&spider=0"

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
	request.Header.Add("User-Agent", "Not_A Brand\";v=\"8\", \"Chromium\";v=\"120\", \"Microsoft Edge\";v=\"120")
	request.Header.Add("Cookie", "_tea_utm_cache_2608=undefined; __tea_cookie_tokens_2608=%257B%2522web_id%2522%253A%25227263387509600912950%2522%252C%2522user_unique_id%2522%253A%25227263387509600912950%2522%252C%2522timestamp%2522%253A1691139204499%257D; uid_tt=d9129948ced9f65f755090d09acb09b9; uid_tt_ss=d9129948ced9f65f755090d09acb09b9; sid_tt=410d35bde6e42e93cc82a7b7d86cad0b; sessionid=410d35bde6e42e93cc82a7b7d86cad0b; sessionid_ss=410d35bde6e42e93cc82a7b7d86cad0b; sid_guard=410d35bde6e42e93cc82a7b7d86cad0b%7C1693644523%7C31535999%7CSun%2C+01-Sep-2024+08%3A48%3A42+GMT; sid_ucp_v1=1.0.0-KGE2ZjRmMDYwNWI5ODE4Yzk1ZWUyYzgxZGQ2NDY3ZGNmMzY5YWI1NDMKFgjoq_CLrc3gAhDr7cunBhiwFDgIQAsaAmxmIiA0MTBkMzViZGU2ZTQyZTkzY2M4MmE3YjdkODZjYWQwYg; ssid_ucp_v1=1.0.0-KGE2ZjRmMDYwNWI5ODE4Yzk1ZWUyYzgxZGQ2NDY3ZGNmMzY5YWI1NDMKFgjoq_CLrc3gAhDr7cunBhiwFDgIQAsaAmxmIiA0MTBkMzViZGU2ZTQyZTkzY2M4MmE3YjdkODZjYWQwYg; store-region=cn-ha; store-region-src=uid; csrf_session_id=a1ad739a0df8eb5fe9da982694aff9be; msToken=6hijmTqCkBVaeUeDp-YNfKGrJFKsNtkCIQFytMyj20fAWZcaovIZgjh1SKrttjR8pBCQh_c8rARGVcDzscbzlqNoThNENwjMVMuD6Rt3gY7-2vGrTx20Uqb4ZVLNimqr")
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
	realtimeData := response.Data

	// 打印解析后的 realtime 数据
	//fmt.Printf("Realtime data: %+v\n", realtimeData)
	data := make([]*Juejin, 0)
	now := time.Now().Unix()
	head := "https://juejin.cn/post/"
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
	}
	for _, item := range data {
		fmt.Println(item)
		// 打印其他字段值...
	}
	model.Conn.Create(data)
}
