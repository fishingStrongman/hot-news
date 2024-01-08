package bilibili

import (
	"encoding/json"
	"fmt"
	"os"
)

//api: https://api.bilibili.com/x/web-interface/wbi/search/square?limit=20&platform=web&w_rid=18e2eb4d6b330093545209bee0f28408&wts=1704525897

const api = "https://api.bilibili.com/x/web-interface/wbi/search/square?limit=20&platform=web"

func Run() {
	getInfo()
}

func getInfo() {
	//client := &http.Client{}
	//request, err := http.NewRequest("GET", api, nil)
	//if err != nil {
	//	fmt.Println("Error creating request:", err)
	//	return
	//}
	//
	//request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	//response, err := client.Do(request)
	//if err != nil {
	//	fmt.Println("Error sending request:", err)
	//	return
	//}
	//
	//defer func() {
	//	_ = response.Body.Close()
	//}()
	//
	//body, err := io.ReadAll(response.Body)
	//if err != nil {
	//	fmt.Println("Error reading response body:", err)
	//	return
	//}
	//fmt.Printf("body:%s", string(body))

	body, _ := os.ReadFile("./app/schedule/bilibili/data.json")
	fmt.Printf("body:%s\n", body)
	var ret Ret
	_ = json.Unmarshal(body, &ret)
	fmt.Printf("ret:%v", ret)
}
