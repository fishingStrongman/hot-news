package baidu

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"hotinfo/app/model"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

const api = "https://top.baidu.com/board?platform=pc&sa=pcindex_entry"

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
	// 发起 GET 请求
	response, err := http.Get(api)
	if err != nil {
		fmt.Println("HTTP request failed:", err)
		return
	}
	defer response.Body.Close()
	// 读取响应内容
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
		return
	}
	// 将 HTML 字符串加载到 GoQuery 文档中
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
	if err != nil {
		fmt.Println("Failed to load HTML:", err)
		return
	}
	//var res []string
	// 提取指定 id 的 div 中的内容
	doc.Find("#sanRoot").Each(func(i int, s *goquery.Selection) {
		divContent, _ := s.Html()
		res := strings.Split(divContent, "热搜榜")
		re := strings.Split(res[0], "appUrl")
		var data []BaiDu
		for i1, s2 := range re {

			if i1 != 0 {
				//获取url
				ss := strings.Split(s2, "&amp")[0]
				u1 := strings.Split(ss, "https://www.baidu.com/")[1]
				url1 := fmt.Sprintf("https://www.baidu.com/%s", u1)
				var titl1 string
				if i1 == 1 {
					//获取标题
					ss1 := strings.Split(s2, "word\":\"")[1]
					a1 := strings.Split(ss1, "\"}")[0]
					titl1 = strings.Split(a1, "\",\"isTop")[0]
				} else {
					//获取标题
					ss1 := strings.Split(s2, "word\":\"")[1]
					titl1 = strings.Split(ss1, "\"}")[0]
				}
				//获取热度
				ss2 := strings.Split(s2, "hotScore\":\"")[1]
				hot1 := strings.Split(ss2, "\",\"hotTag")[0]

				a := BaiDu{
					UpdateVer:   time.Now().Unix(),
					Title:       titl1,
					Url:         url1,
					Hot:         hot1,
					CreatedTime: time.Now().Format("2006-01-02 15:04:05"),
					UpdatedTime: time.Now().Format("2006-01-02 15:04:05"),
				}
				data = append(data, a)
			}

		}
		for _, datum := range data {
			fmt.Println(datum)
		}
		model.Conn.Table(data[0].TableName()).Create(&data)
	})

}
func Refresh() []BaiDu {
	var maxUpdateVer int64

	// 查询最大的 update_ver
	result := model.Conn.Model(&BaiDu{}).Select("MAX(update_ver) as max_update_ver").Scan(&maxUpdateVer)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	// 查询所有 update_ver 为最大值的记录
	var baiduList []BaiDu
	result = model.Conn.Where("update_ver = ?", maxUpdateVer).Find(&baiduList)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	// 打印查询结果
	fmt.Printf("Data with max update_ver (%d):\n", maxUpdateVer)
	return baiduList

}

//func getInfo() {
//	// 创建HTTP客户端
//	client := http.Client{}
//
//	// 创建GET请求
//	request, err := http.NewRequest("GET", api, nil)
//	if err != nil {
//		fmt.Println("Failed to create HTTP request:", err)
//		return
//	}
//
//	// 发送请求并获取响应
//	response, err := client.Do(request)
//	if err != nil {
//		fmt.Println("HTTP request failed:", err)
//		return
//	}
//	defer response.Body.Close()
//
//	// 读取响应内容
//	body, err := ioutil.ReadAll(response.Body)
//	if err != nil {
//		fmt.Println("Failed to read response body:", err)
//		return
//	}
//
//	// 解析响应内容到结构体
//	var baiduHot Hot
//	err = json.Unmarshal(body, &baiduHot)
//	if err != nil {
//		fmt.Println("Failed to unmarshal response body:", err)
//		return
//	}
//	var data []BaiDu
//	now := time.Now().Unix()
//	for _, datum := range baiduHot.Data.List {
//		a := BaiDu{
//			UpdateVer:   now,
//			Title:       datum.Title,
//			Url:         datum.Url,
//			Hot:         datum.Hot,
//			CreatedTime: time.Now().Format("2006-01-02 15:04:05"),
//			UpdatedTime: time.Now().Format("2006-01-02 15:04:05"),
//		}
//		data = append(data, a)
//	}
//	model.Conn.Table(data[0].TableName()).Create(&data)
//}
