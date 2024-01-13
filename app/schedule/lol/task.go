package lol

import (
	"encoding/json"
	"fmt"
	"hotinfo/app/model"
	"io"
	"log"
	"net/http"
	"regexp"
	"time"
)

const api = "https://apps.game.qq.com/cmc/zmMcnTargetContentList?r0=jsonp&page=1&num=16&target=24&source=web_pc&r1=jQuery19108354930441080934_1704804548069&_=1704804548070"

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
func extractJSONP(jsonp []byte) ([]byte, error) {
	re := regexp.MustCompile(`^[^(]*\((.*)\);$`)
	matches := re.FindSubmatch(jsonp)
	if len(matches) < 2 {
		return nil, fmt.Errorf("解析json失败")
	}
	return matches[1], nil
}
func getInfo() {
	client := &http.Client{}
	request, err := http.NewRequest("GET", api, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	request.Header.Add("Cookie", "RK=mpPIt1oMR7; ptcz=817d17a09c36755d9b6b9e4df893f7bda228e409b6b250ee1e3a536123aacdba; pgv_pvid=6835227561; pac_uid=0_a640ac7df3ca5; iip=0; qq_domain_video_guid_verify=54c4d2c196964f47; o_cookie=1790587200; Qs_lvt_323937=1698482180; Qs_pv_323937=2491088095741199000; fqm_pvqid=a426d807-bec2-4830-a33c-e981c2b6d9c1; _clck=1ikym9v|1|fh5|0; eas_sid=91R7W0j2L475Z9g1F7G7j0k4b2; _qimei_uuid42=181090e231d1009be1873878a0aa6f3d6122d55369; _qimei_fingerprint=5c19576b49cd88cdf62f5665aa8d4476; _qimei_q36=; _qimei_h38=e147844c05e0d48c447986ab0200000ae17916; LW_uid=7167o06448B0J4K494I8I0F8p6; isHostDate=19731; PTTuserFirstTime=1704758400000; isOsSysDate=19731; PTTosSysFirstTime=1704758400000; isOsDate=19731; PTTosFirstTime=1704758400000; ts_refer=lol.qq.com/; ts_uid=459324438; weekloop=0-0-0-2; LW_sid=l1S730R43840G4M8Q9P2k3D7S0; pgv_info=ssid=s8077240133; lolqqcomrouteLine=news_index-tool_main_news_index-tool_main_index-tool; tokenParams=%3Fdocid%3D1719755674302751220")
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
	// 解析 JSON 数据
	jsonData, err := extractJSONP(body)
	if err != nil {
		fmt.Println("数据解析失败", err)
	}
	var ret Ret
	// 使用 Unmarshal 解析 JSON 数据到结构体
	err = json.Unmarshal(jsonData, &ret)
	if err != nil {
		fmt.Println("数据读取失败", err)
	}
	data := make([]*Lol, 0)
	// 提取 iDocID 和 sTitle
	for _, result := range ret.Data.Result {
		url := "https://lol.qq.com/news/detail.shtml?docid=" + result.IDocId
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
		data = append(data, &tmp)
	}
	model.Conn.Create(data)
}
func Refresh() []Lol {
	var maxUpdateVer int64

	// 查询最大的 update_ver
	result := model.Conn.Model(&Lol{}).Select("MAX(update_ver) as max_update_ver").Scan(&maxUpdateVer)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	// 查询所有 update_ver 为最大值的记录
	var lolList []Lol
	result = model.Conn.Where("update_ver = ?", maxUpdateVer).Find(&lolList)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	// 打印查询结果
	fmt.Printf("Data with max update_ver (%d):\n", maxUpdateVer)
	return lolList

}
