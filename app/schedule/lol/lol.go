package lol

import (
	"encoding/json"
	"fmt"
	"hotinfo/app/model"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"time"
)

func extractJSONP(jsonp []byte) ([]byte, error) {
	re := regexp.MustCompile(`^[^(]*\((.*)\);$`)
	matches := re.FindSubmatch(jsonp)
	if len(matches) < 2 {
		return nil, fmt.Errorf("解析json失败")
	}
	return matches[1], nil
}

func Run() {
	getLol()
}
func getLol() {
	yurl := "https://apps.game.qq.com/cmc/zmMcnTargetContentList?r0=jsonp&page=1&num=16&target=24&source=web_pc&r1=jQuery19108354930441080934_1704804548069&_=1704804548070"
	// 发送 HTTP 请求
	resp, err := http.Get(yurl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// 解析 JSON 数据
	jsonData, err := extractJSONP(body)
	if err != nil {
		log.Fatal(err)
	}

	// 定义嵌套结构以匹配 JSON 数据
	var responseStruct struct {
		Data struct {
			Result []struct {
				IDocID   string `json:"iDocID"`
				STitle   string `json:"sTitle"`
				SIdxTime string `json:"sIdxTime"`
			} `json:"result"`
		} `json:"data"`
	}

	// 使用 Unmarshal 解析 JSON 数据到结构体
	err = json.Unmarshal(jsonData, &responseStruct)
	if err != nil {
		log.Fatal(err)
	}

	// 提取 iDocID 和 sTitle
	for _, result := range responseStruct.Data.Result {
		url := "https://lol.qq.com/news/detail.shtml?docid=" + result.IDocID
		sTitle := result.STitle
		sIdxTime := result.SIdxTime
		tmp := Lol{
			UpdateVer:   time.Now().Unix(),
			Title:       sTitle,
			Url:         url,
			Time:        sIdxTime,
			CreatedTime: time.Now(),
			UpdatedTime: time.Now(),
		}
		model.Conn.Create(&tmp)
	}
}
